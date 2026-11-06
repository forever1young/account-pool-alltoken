package reporter

import (
	"log"
	"time"
)

// Reporter periodically pushes account status (alive/banned, refresh
// results, usage) back to account-onboard's export API. This is how
// account-onboard's data stays current without ever polling the pool itself.
type Reporter struct {
	baseURL  string
	apiKey   string
	interval time.Duration
}

func New(baseURL, apiKey string, interval time.Duration) *Reporter {
	return &Reporter{baseURL: baseURL, apiKey: apiKey, interval: interval}
}

// Run blocks, reporting on each interval tick.
func (r *Reporter) Run() {
	ticker := time.NewTicker(r.interval)
	defer ticker.Stop()
	for range ticker.C {
		if err := r.reportOnce(); err != nil {
			log.Printf("reporter: report failed: %v", err)
		}
	}
}

func (r *Reporter) reportOnce() error {
	log.Printf("reporter: reportOnce not implemented yet")
	return nil
}
