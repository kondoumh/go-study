package vault

import (
	"bytes"
	"io"
	"time"
)

type secret struct {
	ID         string
	CreateTime time.Time
	token      string
}

func (s *secret) Read(p []byte) (int, error) {
	return bytes.NewBuffer(p).WriteString(s.token)
}

func NewSecret() io.Reader {
	return &secret{
		ID:         "dummy_id",
		CreateTime: time.Now(),
		token:      "dummy_token",
	}
}
