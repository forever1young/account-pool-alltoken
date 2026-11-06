package service

// FingerprintApplier applies the TLS/CC fingerprint parameters copied from
// account-onboard at pull time to outgoing Claude.ai requests, avoiding bot detection.
type FingerprintApplier struct{}

func NewFingerprintApplier() *FingerprintApplier {
	return &FingerprintApplier{}
}

func (f *FingerprintApplier) Apply(accountID uint) error {
	panic("not implemented")
}
