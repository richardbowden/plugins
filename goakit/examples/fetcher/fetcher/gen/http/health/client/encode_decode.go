// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// health HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/fetcher/fetcher/design

package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/http"
)

// BuildShowRequest instantiates a HTTP request object with method and path set
// to call the "health" service "show" endpoint
func (c *Client) BuildShowRequest(v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowHealthPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("health", "show", u.String(), err)
	}

	return req, nil
}

// DecodeShowResponse returns a decoder for responses returned by the health
// show endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("health", "show", err)
			}

			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("account", "create", resp.StatusCode, string(body))
		}
	}
}