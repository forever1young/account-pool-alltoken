package service

// ProxyPool binds each account to a dedicated outbound proxy IP so every
// request to Claude.ai for that account exits from the same address.
type ProxyPool struct{}

func NewProxyPool() *ProxyPool {
	return &ProxyPool{}
}

func (p *ProxyPool) ProxyForAccount(accountID uint) (address string, err error) {
	panic("not implemented")
}

func (p *ProxyPool) AssignProxy(accountID uint, region string) error {
	panic("not implemented")
}
