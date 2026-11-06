# account-pool 开发清单

当前状态：仅有骨架，service 方法多为 `panic("not implemented")`，
controller 返回 501，puller/reporter 的 `Run()` 已有定时循环但 `pullOnce`/`reportOnce` 是空实现。

设计依据与接口契约在上号平台仓库：
`account-onboard-alltoken/docs/architecture.md` 和 `docs/export-api.md`。

## 阶段一：基础设施

- [ ] **数据库连接** — `model/db.go`：GORM + postgres driver，读 `DATABASE_URL`，
      `AutoMigrate`（Account / Proxy / UsageLog / ClientAPIKey）
- [ ] **Redis 连接** — `go-redis`，读 `REDIS_URL`
- [ ] **Dockerfile** — `docker-compose.yml` 写了 `build: .` 但缺 Dockerfile，需补多阶段构建
- [ ] **启动时配置校验** — `DATABASE_URL`、`CREDENTIAL_ENCRYPTION_KEY`、
      `ONBOARD_EXPORT_BASE_URL`、`ONBOARD_EXPORT_API_KEY` 为空直接退出
      （现有 claude-gateway 镜像就是因为连接串缺失/格式错，只在运行时报
      `connection URL requires a postgres or postgresql scheme and host` 这种难定位的错）

## 阶段二：凭据加密

- [ ] **`service/crypto.go`** — AES-256-GCM，和平台**不同**的 key
  - 从平台拉来的明文 token 要用本服务的 key 重新加密入库
  - key 一旦有账号入库不可更换

## 阶段三：账号拉取（接入平台，核心）

- [ ] **`service/puller/pullOnce`**
  - HTTP GET `{ONBOARD_EXPORT_BASE_URL}/api/export/accounts?cursor=&limit=`
  - 请求头 `X-Export-Api-Key: {ONBOARD_EXPORT_API_KEY}`
  - 响应里的 token 用本服务 key 加密，按 `OnboardAccountID` upsert（唯一索引已建好）
  - 入库成功后 POST `/api/export/ack` 回传 account_ids
  - cursor 持久化（存 Redis 或 DB），重启后从上次位置继续
  - **幂等**：ack 失败时下次会重复拉到同一批，upsert 必须能安全重复执行
  - 失败只记日志不中断循环（`Run()` 已经这样写了）
- [ ] **拉取时同步指纹参数** — 响应里的 `fingerprint_template.params` 存进
      `Account.FingerprintParams`
- [ ] **代理分配** — 新账号入库后按 `proxy_region_hint` 调 ProxyPool.AssignProxy

## 阶段四：状态上报（平台的数据更新机制）

- [ ] **`service/reporter/reportOnce`**
  - 收集账号状态：status（active/cooling_down/banned）、last_used_at、
    token 刷新结果、用量摘要
  - POST `{ONBOARD_EXPORT_BASE_URL}/api/export/report`，同样带 API Key
  - 建议只上报自上次以来有变化的账号（增量），全量上报要控制批次大小

## 阶段五：Token 刷新

- [ ] **`service/token_refresh.go RefreshIfNeeded`**
  - access_token 快过期时用 refresh_token 换新的
  - 刷新请求也要走该账号绑定的代理 IP（和模型请求同一出口，避免 IP 跳变）
  - 刷新失败（refresh_token 失效/账号被封）→ 账号状态改 `banned`，由 reporter 上报平台
- [ ] **定时刷新任务** — 单独 goroutine 扫描快过期的账号，`main.go` 里和
      puller/reporter 一起起

## 阶段六：代理 IP 池

- [ ] **`service/proxy_pool.go`**
  - 对接代理服务商 API（`PROXY_POOL_PROVIDER_URL`/`_KEY`）获取/续期 IP
  - `AssignProxy` — 按地区给账号分配独立出口 IP，写入 `Account.ProxyID`
  - `ProxyForAccount` — 返回该账号的代理地址，供 relay 使用
  - 同一账号必须始终用同一出口 IP（IP 跳变会触发 Claude.ai 风控）
  - IP 失效时的替换策略

## 阶段七：指纹伪装

- [ ] **`service/fingerprint.go Apply`**
  - TLS 指纹伪装，Go 侧一般用 `utls`（uTLS）替换标准库 TLS 握手
  - 应用 `Account.FingerprintParams` 里的 CC/TLS 参数
  - 和 relay 的 HTTP 客户端打通：每个账号的请求用自己的指纹 + 自己的代理

## 阶段八：调度

- [ ] **`service/scheduler.go PickAccount`**
  - 按模型筛选可用账号（status=active、token 有效）
  - 并发控制：每账号最大并发（Redis 计数器），满了选下一个
  - 冷却：触发限流/报错的账号临时降权（status=cooling_down + TTL）
  - 负载均衡：轮询或最少使用优先
  - 无可用账号时返回明确错误，别让请求挂死

## 阶段九：转发（对外服务的核心）

- [ ] **`relay/relay.go ForwardChatCompletion`**
  - OpenAI 格式请求 → Claude 原生格式（messages、system、max_tokens 等映射）
  - 用 scheduler 选中账号的凭据 + 代理 + 指纹发请求
  - Claude 响应 → OpenAI 格式响应
  - **流式响应（SSE）**：Claude 的 event stream 转成 OpenAI 的 `data: {...}` chunk 格式，
    这块最容易出 bug，注意首包延迟和 `[DONE]` 结尾
  - 上游报错的处理：401/403 → 账号可能失效，标记并换账号重试；
    429 → 账号冷却；其他 5xx → 有限重试
- [ ] **`controller/openai_compat.go ChatCompletions`** — 对接 relay，支持 stream 参数
- [ ] **`controller/openai_compat.go ListModels`** — 返回支持的模型列表

## 阶段十：客户端鉴权与计量

- [ ] **`middleware/client_auth.go` 改造** — 目前只比对单个环境变量 `CLIENT_API_KEY`，
      需要改成查 `ClientAPIKey` 表（KeyHash 已有唯一索引），支持多客户端多 key
- [ ] **限流** — 按客户端 key 限流（Redis）
- [ ] **用量记账** — 每次请求写 `UsageLog`（account_id、client_key_id、model、token 数）
      ⚠️ token 数做上界校验，别让上游返回的异常大数值直接入库

## 阶段十一：管理前端（`web/`）

- [ ] 账号池状态看板：各账号 status、绑定的出口 IP、最近使用、token 有效期
- [ ] 用量统计：按账号/按客户端/按模型
- [ ] 客户端 API Key 管理：新建、吊销
- [ ] 代理池状态

## 阶段十二：上线准备

- [ ] 号池出口 IP 固定下来，提供给平台加白名单
- [ ] 对 Claude.ai 的全部流量确认都走了代理（不要有漏网的直连）
- [ ] `.env` 备份（`CREDENTIAL_ENCRYPTION_KEY` 丢了等于所有账号凭据作废）
- [ ] 日志轮转（`--log-opt max-size`），号池日志量大

## 约束

- ✅ 所有和平台之间的连接都由本服务发起（pull / ack / report），
  平台永远不会主动连本服务
- ✅ 本服务的加密 key 必须和平台的不同
