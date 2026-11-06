package puller

import (
	"log"
	"time"
)

// Puller periodically pulls newly delivered accounts from account-onboard's
// export API and acknowledges receipt. This is the only outbound connection
// account-pool makes toward account-onboard for new-account intake.
type Puller struct {
	baseURL  string
	apiKey   string
	interval time.Duration
}

func New(baseURL, apiKey string, interval time.Duration) *Puller {
	return &Puller{baseURL: baseURL, apiKey: apiKey, interval: interval}
}

// Run blocks, pulling on each interval tick.
func (p *Puller) Run() {
	ticker := time.NewTicker(p.interval)
	defer ticker.Stop()
	for range ticker.C {
		if err := p.pullOnce(); err != nil {
			log.Printf("puller: pull failed: %v", err)
		}
	}
}

func (p *Puller) pullOnce() error {
	log.Printf("puller: pullOnce not implemented yet")
	return nil
}
