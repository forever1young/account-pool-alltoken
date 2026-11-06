package service

// TokenRefresher keeps account access tokens fresh using their refresh_token,
// independent of account-onboard once an account has been pulled.
type TokenRefresher struct{}

func NewTokenRefresher() *TokenRefresher {
	return &TokenRefresher{}
}

func (t *TokenRefresher) RefreshIfNeeded(accountID uint) error {
	panic("not implemented")
}
