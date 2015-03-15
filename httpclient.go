package interfaces

import (
	"io"
	"net/http"
	"net/url"
)

// Interface maps to http.Client
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
	Get(string) (*http.Response, error)
	Head(string) (*http.Response, error)
	Post(string, string, io.Reader) (*http.Response, error)
	PostForm(string, url.Values) (*http.Response, error)
}
