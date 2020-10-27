package productparser

import (
	"context"
	"encoding/csv"
	"io"
	"strconv"
	"strings"

	"github.com/isavinof/pricer/lib/log"
	"github.com/isavinof/pricer/lib/types"
	"github.com/pkg/errors"
)

// CsvParser parse csv data
type CsvParser struct {
}

// NewCsvParser object initialization only
func NewCsvParser() *CsvParser {
	return &CsvParser{}
}

// AcceptedFormat provide allowed formats in HTTP style
func (parser *CsvParser) AcceptedFormat() string {
	return "text/csv; charset=utf-8"
}

// ParseFromReader
func (parser *CsvParser) ParseFromReader(ctx context.Context, reader io.Reader) ([]types.ProductPrice, error) {
	logger := log.FromContext(ctx)
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.Wrapf(err, "read data from responded file")
	}

	if len(records) == 0 {
		return nil, nil
	}

	products := make([]types.ProductPrice, 0, len(records))
	logger.Infof("start parse data from")
	for _, record := range records {
		if len(record) < 2 {
			// here could be continue without error for all function. It depends on requirements
			return nil, errors.Wrapf(err, "unexpected record format: %v. Expected: $name;$price", strings.Join(record, string(csvReader.Comma)))
		}
		price, err := strconv.ParseFloat(record[1], 32)
		if err != nil {
			// here could be continue without error for all function. It depends on requirements
			return nil, errors.Wrapf(err, "unexpected price format: %v. Expected: float", record[1])
		}
		products = append(products, types.ProductPrice{
			ProductName:       record[0],
			ProductPriceCents: int32(price * 100),
		})
	}

	return products, nil
}
