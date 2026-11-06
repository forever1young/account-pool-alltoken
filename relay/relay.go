package relay

// Relay forwards a client request to Claude.ai using a scheduled account's
// credentials and proxy, converting between the client's OpenAI-compatible
// format and Claude's native format, including streaming responses.
type Relay struct{}

func New() *Relay {
	return &Relay{}
}

func (r *Relay) ForwardChatCompletion(accountID uint, body []byte) ([]byte, error) {
	panic("not implemented")
}
