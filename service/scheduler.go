package service

// Scheduler picks an eligible account for an incoming request, accounting
// for concurrency limits, cooldowns, and load balancing across the pool.
type Scheduler struct{}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) PickAccount(model string) (accountID uint, err error) {
	panic("not implemented")
}
