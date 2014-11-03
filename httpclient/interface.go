package httpclient

import (
	"net/http"
)

type Interface interface {
	Do(*http.Request) (*http.Response, error)
}
