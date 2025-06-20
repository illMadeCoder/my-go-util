package mockstream

import (
	"io"
	"math/rand"
	"time"
)

type MockStream struct {
	rand      *rand.Rand
	remaining int
	written   []byte
}

type MockStreamConfig struct {
	Seed   int64
	Length int
}

func NewMockStream(cfg MockStreamConfig) *MockStream {
	rndsrc := rand.NewSource(cfg.Seed)
	r := rand.New(rndsrc)
	return &MockStream{
		rand:      r,
		remaining: cfg.Length,
	}
}

func (m *MockStream) Write(b []byte) (n int, err error) {
	time.Sleep(time.Second)
	m.written = append(m.written, b...)
	return len(b), nil
}

func (m *MockStream) Read(b []byte) (n int, err error) {
	if m.remaining == 0 {
		return 0, io.EOF
	}

	for i := range b {
		b[i] = byte(m.rand.Intn(95) + 32)
		m.remaining -= 1
		if m.remaining == 0 {
			return i + 1, nil
		}
	}
	return len(b), nil
}
