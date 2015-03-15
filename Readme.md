# interfaces

[![GoDoc](https://godoc.org/gopkg.in/httgo/interfaces.v3?status.svg)](http://godoc.org/gopkg.in/httgo/interfaces.v3)

`net/http` Client interface

## Install

    go get gopkg.in/httgo/interfaces.v3

## Usage

    type ApiClient struct {
      C interfaces.HTTPClient
    }

    func NewClient(opts ...func(c *ApiClient)) *ApiClient {
      c := &ApiClient{
        C: http.DefaultClient,
      }
      for _, v := range opts {
        v(c)
      }
      return c
    }

In your tests you can;

    api := NewClient(func(c *ApiClient) {
      c.C = &MyMockHttpClient{}
    })

## License

MIT

