package main

import (
	"log"
	"time"

	"github.com/forever1young/account-pool-alltoken/config"
	"github.com/forever1young/account-pool-alltoken/router"
	"github.com/forever1young/account-pool-alltoken/service/puller"
	"github.com/forever1young/account-pool-alltoken/service/reporter"
)

func main() {
	cfg := config.Load()

	go puller.New(cfg.OnboardExportBaseURL, cfg.OnboardExportAPIKey,
		time.Duration(cfg.OnboardPullIntervalSeconds)*time.Second).Run()
	go reporter.New(cfg.OnboardExportBaseURL, cfg.OnboardExportAPIKey,
		time.Duration(cfg.OnboardReportIntervalSeconds)*time.Second).Run()

	r := router.New(cfg)

	log.Printf("account-pool listening on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
