package httpclient

import (
	"net/http"
)

// Interface maps to http.Client
type Interface interface {
	Do(*http.Request) (*http.Response, error)
}
