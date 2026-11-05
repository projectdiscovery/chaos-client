package runner

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
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

func TestWriteSubdomainWithOutputFilePrintsOnce(t *testing.T) {
	dir := t.TempDir()
	stdout, err := os.Create(filepath.Join(dir, "stdout"))
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = stdout.Close() }()
	origStdout := os.Stdout
	os.Stdout = stdout
	defer func() { os.Stdout = origStdout }()

	outputFile, err := os.Create(filepath.Join(dir, "out.txt"))
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = outputFile.Close() }()

	// same writers RunEnumeration sets up for -o
	opts := &Options{Domain: "example.com", Output: outputFile.Name(), outputFile: outputFile}
	opts.outputWriter = io.MultiWriter(os.Stdout, outputFile)

	if err := writeSubdomain(opts, "www"); err != nil {
		t.Fatal(err)
	}

	for _, f := range []string{stdout.Name(), outputFile.Name()} {
		data, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		if n := strings.Count(string(data), "www.example.com"); n != 1 {
			t.Fatalf("%s has %d copies of www.example.com, want 1: %q", filepath.Base(f), n, data)
		}
	}
}
