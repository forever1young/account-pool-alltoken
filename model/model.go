package model

import "time"

// Account is a Claude.ai account pulled from account-onboard. Credential
// fields are re-encrypted with this service's own CredentialEncryptionKey,
// independent of account-onboard's key.
type Account struct {
	ID                    uint `gorm:"primaryKey"`
	OnboardAccountID      uint `gorm:"uniqueIndex"`
	OrganizationID        string `gorm:"index"`
	Email                 string
	EncryptedRefreshToken string
	EncryptedAccessToken  string
	AccessTokenExpiresAt  time.Time
	FingerprintParams     string // JSON blob, copied from onboard at pull time
	ProxyID               *uint
	Status                string // active | cooling_down | banned
	LastUsedAt            *time.Time
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

// Proxy is an outbound IP bound to one or more accounts.
type Proxy struct {
	ID        uint `gorm:"primaryKey"`
	Address   string
	Region    string
	CreatedAt time.Time
}

// UsageLog records per-request billing/accounting data.
type UsageLog struct {
	ID         uint `gorm:"primaryKey"`
	AccountID  uint
	ClientKeyID uint
	Model      string
	PromptTokens     int
	CompletionTokens int
	CreatedAt  time.Time
}

// ClientAPIKey authenticates business clients against the OpenAI-compatible API.
type ClientAPIKey struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	KeyHash   string `gorm:"uniqueIndex"`
	CreatedAt time.Time
}
