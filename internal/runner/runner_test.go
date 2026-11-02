package runner

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriteJSONLineOneNewline(t *testing.T) {
	payload := `{"domain":"example.com","subdomains":["sub1","sub2"]}`
	for _, body := range []string{payload, payload + "\n", payload + "\r\n"} {
		buf := &bytes.Buffer{}
		writeJSONLine(buf, strings.NewReader(body))
		if buf.String() != payload+"\n" {
			t.Fatalf("body %q produced %q", body, buf.String())
		}
	}
}
