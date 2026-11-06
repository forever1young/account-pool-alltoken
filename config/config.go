package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port                        string
	DatabaseURL                 string
	RedisURL                    string
	CredentialEncryptionKey     string
	OnboardExportBaseURL        string
	OnboardExportAPIKey         string
	OnboardPullIntervalSeconds  int
	OnboardReportIntervalSeconds int
	ClientAPIKey                string
	ProxyPoolProviderURL        string
	ProxyPoolProviderKey        string
}

func Load() *Config {
	return &Config{
		Port:                         envOr("PORT", "8787"),
		DatabaseURL:                  os.Getenv("DATABASE_URL"),
		RedisURL:                     os.Getenv("REDIS_URL"),
		CredentialEncryptionKey:      os.Getenv("CREDENTIAL_ENCRYPTION_KEY"),
		OnboardExportBaseURL:         os.Getenv("ONBOARD_EXPORT_BASE_URL"),
		OnboardExportAPIKey:          os.Getenv("ONBOARD_EXPORT_API_KEY"),
		OnboardPullIntervalSeconds:   envIntOr("ONBOARD_PULL_INTERVAL_SECONDS", 30),
		OnboardReportIntervalSeconds: envIntOr("ONBOARD_REPORT_INTERVAL_SECONDS", 60),
		ClientAPIKey:                 os.Getenv("CLIENT_API_KEY"),
		ProxyPoolProviderURL:         os.Getenv("PROXY_POOL_PROVIDER_URL"),
		ProxyPoolProviderKey:         os.Getenv("PROXY_POOL_PROVIDER_KEY"),
	}
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envIntOr(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}
