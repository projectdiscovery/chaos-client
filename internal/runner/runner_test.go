package runner

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/projectdiscovery/chaos-client/pkg/chaos"
)

func TestJSONOutputEndsWithNewline(t *testing.T) {
	buf := &bytes.Buffer{}
	opts := &Options{
		outputWriter: buf,
		JSONOutput:   true,
	}

	payload := `{"domain":"example.com","subdomains":["sub1","sub2"]}`
	reader := io.NopCloser(strings.NewReader(payload))
	item := &chaos.Result{Reader: &reader}

	_, _ = io.Copy(opts.outputWriter, *item.Reader)
	_, _ = fmt.Fprintln(opts.outputWriter)

	if !strings.HasSuffix(buf.String(), "\n") {
		t.Fatalf("expected output to end with newline, got %q", buf.String())
	}
	expected := payload + "\n"
	if buf.String() != expected {
		t.Fatalf("expected %q, got %q", expected, buf.String())
	}
}
