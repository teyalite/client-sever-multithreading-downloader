package network

import (
	"net"
	"net/http"
	"net/url"
	"time"
)

// Build request object with custom header.
func BuildRequest(method string, url string) (*http.Request, error) {
	r, err := http.NewRequest(method, url, nil)
	// sumulating a browser user-agent
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36")
	return r, err
}

/*
 * HTTP Client optimized for concurrent download
 * https://stackoverflow.com/questions/41719797/tls-handshake-timeout-on-requesting-data-concurrently-from-api
 */
func HTTPClient() *http.Client {
	t := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 600 * time.Second,
	}
	return &http.Client{Transport: t}
}

// URL validation.
func IsValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
