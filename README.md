# account-pool

号池：托管 Claude.ai 账号、按账号绑定代理出口 IP，对外提供 OpenAI 兼容 API。

账号来源是 [account-onboard](https://github.com/forever1young/account-onboard-alltoken)：
本服务定时主动拉取新账号、确认交付、再定时上报账号状态。所有连接均由本服务发起，
account-onboard 从不主动连接本服务。导出 API 契约见 account-onboard 仓库的
`docs/export-api.md`。

## 结构

```
cmd/server/        — 入口
config/            — 环境变量加载
router/            — 路由注册
controller/
  openai_compat.go — 对外 /v1 OpenAI 兼容 API
  admin.go         — 号池管理界面接口
service/
  puller/          — 定时拉取 account-onboard 导出的新账号 + ack
  reporter/        — 定时上报账号状态给 account-onboard
  token_refresh.go — refresh_token 定时刷新
  scheduler.go     — 账号选取 / 并发 / 降权调度
  proxy_pool.go    — 代理 IP 池，账号与出口 IP 绑定
  fingerprint.go   — TLS/CC 指纹应用
  crypto.go        — 凭据加密（号池自有密钥，与 account-onboard 隔离）
relay/             — 转发 Claude.ai 请求、流式响应、格式转换
model/             — GORM 数据模型
middleware/        — 客户端 API Key 鉴权、限流、并发控制
web/               — 管理前端
migrations/        — 数据库迁移
```

## 开发

```bash
cp .env.example .env
go run ./cmd/server
```
