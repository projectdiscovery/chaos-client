package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAll(t *testing.T) {
	t.Run("Authorization token not specified", func(t *testing.T) {
		os.Args = []string{"", "subdomains", "hackerone.com"}
		err := Run(nil)
		require.EqualError(t, err, "Authorization token not specified")
	})

	t.Run("subdomains", func(t *testing.T) {
		t.Run("API Response 500 error", func(t *testing.T) {
			os.Args = []string{"", "subdomains", "test.com", "-k", "sdf"}

			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(500)
				w.Write([]byte(`err`))
			}))
			defer svr.Close()
			err := Run(&svr.URL)
			require.EqualError(t, err, "could not get subdomains for test.com: invalid status code received: 500 - err")
		})

		t.Run("Stats  Response 200", func(t *testing.T) {
			storeStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			os.Args = []string{"", "subdomains", "test.com", "-k", "sdf", "--count"}

			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte(`{"domain":"test.com","subdomains":2}`))
			}))
			defer svr.Close()
			err := Run(&svr.URL)
			w.Close()
			out, _ := io.ReadAll(r)
			os.Stdout = storeStdout
			require.NoError(t, err)
			exp := `2`

			require.Equal(t, exp, strings.TrimSpace(string(out)))
		})

		t.Run("Stats Response 200 (json)", func(t *testing.T) {
			storeStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			os.Args = []string{"", "subdomains", "test.com", "-k", "sdf", "--count", "--json"}

			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte(`{"domain":"test.com","subdomains":2}`))
			}))
			defer svr.Close()
			err := Run(&svr.URL)
			w.Close()
			os.Stdout = storeStdout
			out, _ := io.ReadAll(r)
			require.NoError(t, err)
			exp := `{"count":2}\n`

			require.Equal(t, exp, strings.TrimSpace(string(out)))
		})

		t.Run("API Response 200", func(t *testing.T) {
			storeStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			os.Args = []string{"", "subdomains", "test.com", "-k", "sdf"}

			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte(`{"domain":"test.com","subdomains": ["a","b","c"],"count":3}`))
			}))
			defer svr.Close()
			err := Run(&svr.URL)
			w.Close()
			out, _ := io.ReadAll(r)
			os.Stdout = storeStdout
			require.NoError(t, err)
			exp := `a.test.com
b.test.com
c.test.com`

			require.Equal(t, exp, strings.TrimSpace(string(out)))
		})

		t.Run("API Response 200 (json)", func(t *testing.T) {
			storeStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			os.Args = []string{"", "subdomains", "test.com", "-k", "sdf", "--json"}

			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte(`{"domain":"test.com","subdomains": ["a","b","c"],"count":3}`))
			}))
			defer svr.Close()
			err := Run(&svr.URL)
			w.Close()
			os.Stdout = storeStdout
			out, _ := io.ReadAll(r)
			require.NoError(t, err)
			exp := `{"domain":"test.com","subdomains": ["a","b","c"],"count":3}`

			require.Equal(t, exp, strings.TrimSpace(string(out)))
		})

		t.Run("API Response 200 (json, silent)", func(t *testing.T) {
			storeStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			os.Args = []string{"", "subdomains", "test.com", "-k", "sdf", "--json", "--silent"}

			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte(`{"domain":"test.com","subdomains": ["a","b","c"],"count":3}`))
			}))
			defer svr.Close()
			err := Run(&svr.URL)
			w.Close()
			os.Stdout = storeStdout
			out, _ := io.ReadAll(r)
			require.NoError(t, err)
			exp := `{"domain":"test.com","subdomains": ["a","b","c"],"count":3}`

			require.Equal(t, exp, strings.TrimSpace(string(out)))
		})

		t.Run("API Response 200 (out to file)", func(t *testing.T) {
			os.Args = []string{"", "subdomains", "test.com", "-k", "sdf", "-o", "out1"}

			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte(`{"domain":"test.com","subdomains": ["a","b","c"],"count":3}`))
			}))
			defer svr.Close()
			err := Run(&svr.URL)
			require.NoError(t, err)
			exp := `a.test.com
b.test.com
c.test.com`

			rdr, err := os.Open("out1")
			require.NoError(t, err)
			res, err := io.ReadAll(rdr)
			require.NoError(t, err)
			require.Equal(t, exp, strings.TrimSpace(string(res)))
		})

		t.Run("API Response 200 (out to file, json)", func(t *testing.T) {
			os.Args = []string{"", "subdomains", "test.com", "-k", "sdf", "-o", "out2", "--json"}

			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte(`{"domain":"test.com","subdomains": ["a","b","c"],"count":3}`))
			}))
			defer svr.Close()
			err := Run(&svr.URL)
			require.NoError(t, err)
			exp := `{"domain":"test.com","subdomains": ["a","b","c"],"count":3}`

			rdr, err := os.Open("out2")
			require.NoError(t, err)
			res, err := io.ReadAll(rdr)
			require.NoError(t, err)
			require.Equal(t, exp, strings.TrimSpace(string(res)))
		})
	})

	t.Run("subdomains-batch", func(t *testing.T) {
		t.Run("file not found", func(t *testing.T) {
			os.Args = []string{"", "subdomains-batch", "test.com", "-k", "sdf"}

			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(500)
				w.Write([]byte(`err`))
			}))
			defer svr.Close()
			err := Run(&svr.URL)
			require.EqualError(t, err, "could not open file test.com: open test.com: no such file or directory")
		})

		t.Run("API Response 200", func(t *testing.T) {
			storeStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			os.Args = []string{"", "subdomains-batch", "domains", "-k", "sdf"}

			i := 0
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				if i == 0 {
					require.Equal(t, "/dns/t1.com/subdomains", r.URL.String())
					w.Write([]byte(`{"domain":"t1.com","subdomains": ["a","b"],"count":3}`))
				} else {
					require.Equal(t, "/dns/t2.com/subdomains", r.URL.String())
					w.Write([]byte(`{"domain":"t2.com","subdomains": ["c","d"],"count":3}`))
				}
				i++
			}))
			defer svr.Close()
			d1 := []byte("t1.com\nt2.com")
			err := os.WriteFile("domains", d1, 0644)
			require.NoError(t, err)
			err = Run(&svr.URL)
			w.Close()
			os.Stdout = storeStdout
			out, _ := io.ReadAll(r)
			require.NoError(t, err)
			exp := `a.t1.com
b.t1.com
c.t2.com
d.t2.com`

			require.Equal(t, exp, strings.TrimSpace(string(out)))
		})
	})

	t.Run("dns", func(t *testing.T) {
		res := `{"a":["a"],"aaaa":["aaaa"],"ttl":123}`
		t.Run("file not found", func(t *testing.T) {
			os.Args = []string{"", "dns", "test.com", "-k", "sdf"}

			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(500)
				w.Write([]byte(`err`))
			}))
			defer svr.Close()
			err := Run(&svr.URL)
			require.EqualError(t, err, "could not lookup domain: could not make request: wrong status 500")
		})

		t.Run("API Response 200", func(t *testing.T) {
			storeStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			os.Args = []string{"", "dns", "domain.com", "-k", "sdf"}

			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte(res))
			}))
			defer svr.Close()
			err := Run(&svr.URL)
			w.Close()
			os.Stdout = storeStdout
			out, _ := io.ReadAll(r)
			require.NoError(t, err)
			exp := `{"a":["a"],"aaaa":["aaaa"],"ttl":123}`

			require.Equal(t, exp, strings.TrimSpace(string(out)))
		})

		t.Run("API Response 200 (A)", func(t *testing.T) {
			storeStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			os.Args = []string{"", "dns", "domain.com", "-k", "sdf", "-t", "a"}

			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.Write([]byte(res))
			}))
			defer svr.Close()
			err := Run(&svr.URL)
			w.Close()
			os.Stdout = storeStdout
			out, _ := io.ReadAll(r)
			require.NoError(t, err)
			exp := `{"a":["a"]}`

			require.Equal(t, exp, strings.TrimSpace(string(out)))
		})
	})
}
