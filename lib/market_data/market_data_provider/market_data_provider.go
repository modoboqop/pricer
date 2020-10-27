package marketdata

import (
	"context"
	"io"

	httprequest "github.com/imroc/req"
	"github.com/isavinof/pricer/lib/log"
	"github.com/isavinof/pricer/lib/types"
	"github.com/pkg/errors"
)

//go:generate mockgen -source=market_data_provider.go -destination=market_data_provider_mock.go -package=marketdata ResponseParser

// RestProvider request csv from url and return products
type RestProvider struct {
	parser ResponseParser
}

// NewRestProvider pass parser which conform ResponseParser interface
func NewRestProvider(parser ResponseParser) *RestProvider {
	return &RestProvider{
		parser: parser,
	}
}

// ResponseParser allow manage accepted formats and change it to json, xml, etc.
type ResponseParser interface {
	ParseFromReader(ctx context.Context, reader io.Reader) ([]types.ProductPrice, error)
	AcceptedFormat() string
}

// Fetch fetch product prices from passed url
func (provider *RestProvider) Fetch(ctx context.Context, url string) ([]types.ProductPrice, error) {
	// TODO: verify URL (allow/declined) before request
	logger := log.FromContext(ctx)

	logger.Infof("start request data from: %v", url)

	headers := httprequest.Header{"Accept": provider.parser.AcceptedFormat()}

	req := httprequest.New()
	resp, err := req.Get(url, headers, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get request from url")
	}

	logger.Infof("end request data from: %v", url)

	defer resp.Response().Body.Close()

	return provider.parser.ParseFromReader(ctx, resp.Response().Body)
}
