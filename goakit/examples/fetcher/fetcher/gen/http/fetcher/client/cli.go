// Code generated by goa v3.0.7, DO NOT EDIT.
//
// fetcher HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/plugins/v3/goakit/examples/fetcher/fetcher/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/fetcher/fetcher

package client

import (
	fetcher "goa.design/plugins/v3/goakit/examples/fetcher/fetcher/gen/fetcher"
)

// BuildFetchPayload builds the payload for the fetcher fetch endpoint from CLI
// flags.
func BuildFetchPayload(fetcherFetchURL string) (*fetcher.FetchPayload, error) {
	var url_ string
	{
		url_ = fetcherFetchURL
	}
	payload := &fetcher.FetchPayload{
		URL: url_,
	}
	return payload, nil
}
