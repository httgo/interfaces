package interfaces

import (
	"net/http"
)

// Interface maps to http.Client
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}
