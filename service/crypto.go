package service

// CredentialCipher encrypts/decrypts Account credential fields at rest using
// this service's own CredentialEncryptionKey — independent of account-onboard's
// key, so the two systems never share a secret.
type CredentialCipher struct {
	key []byte
}

func NewCredentialCipher(key string) *CredentialCipher {
	return &CredentialCipher{key: []byte(key)}
}

func (c *CredentialCipher) Encrypt(plaintext string) (string, error) {
	panic("not implemented")
}

func (c *CredentialCipher) Decrypt(ciphertext string) (string, error) {
	panic("not implemented")
}
