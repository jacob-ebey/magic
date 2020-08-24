package admin

import (
	"io"
	"net/http"
)

type clientMock struct {
	calls []*http.Request
	res   *http.Response
}

func (c *clientMock) Do(req *http.Request) (*http.Response, error) {
	c.calls = append(c.calls, req)

	return c.res, nil
}

type bodyMock struct {
	reader io.Reader
}

func (b *bodyMock) Close() error {
	return nil
}
func (b *bodyMock) Read(p []byte) (n int, err error) {
	if b.reader != nil {
		return b.reader.Read(p)
	}

	return 0, nil
}
