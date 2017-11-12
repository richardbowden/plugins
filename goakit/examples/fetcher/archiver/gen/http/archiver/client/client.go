// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// archiver client HTTP transport
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/fetcher/archiver/design

package client

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	goahttp "goa.design/goa/http"
)

// Client lists the archiver service endpoint HTTP clients.
type Client struct {
	// Archive Doer is the HTTP client used to make requests to the archive
	// endpoint.
	ArchiveDoer goahttp.Doer

	// Read Doer is the HTTP client used to make requests to the read endpoint.
	ReadDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the archiver service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ArchiveDoer:         doer,
		ReadDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Archive returns a endpoint that makes HTTP requests to the archiver service
// archive server.
func (c *Client) Archive() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeArchiveRequest(c.encoder)
		decodeResponse = DecodeArchiveResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildArchiveRequest(v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}

		resp, err := c.ArchiveDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("archiver", "archive", err)
		}
		return decodeResponse(resp)
	}
}

// Read returns a endpoint that makes HTTP requests to the archiver service
// read server.
func (c *Client) Read() endpoint.Endpoint {
	var (
		decodeResponse = DecodeReadResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildReadRequest(v)
		if err != nil {
			return nil, err
		}

		resp, err := c.ReadDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("archiver", "read", err)
		}
		return decodeResponse(resp)
	}
}