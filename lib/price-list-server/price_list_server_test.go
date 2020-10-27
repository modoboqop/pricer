package pricelistserver

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/isavinof/pricer/lib/config"
	"github.com/isavinof/pricer/lib/log"
	pricelist "github.com/isavinof/pricer/lib/price-list"
	"github.com/isavinof/pricer/lib/types"
	"github.com/stretchr/testify/assert"
)

var run = func(ctx context.Context, timeout time.Duration, f func(ctx context.Context) error) error {
	return f(ctx)
}

func TestPriceListServer_List(t *testing.T) {

	t.Run("when store error", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		store := NewMockMarketDataStore(ctrl)

		server := NewPriceListServer(config.ServerConfig{StoreTimeout: time.Second, RestTimeout: time.Hour}, log.NewLogger(), nil, store, run)

		req := &pricelist.ListRequest{
			Limit:            10,
			Offset:           20,
			SortingType:      pricelist.SortingType_SortByProductName,
			SortingDirection: pricelist.SortingDirection_SortAsc,
		}

		store.EXPECT().Get(gomock.Any(), types.SortByProductName, types.SortAsc, int64(10), int64(20)).Return(nil, errors.New("error"))
		got, err := server.List(context.Background(), req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})

	t.Run("when store ok and data empty", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		store := NewMockMarketDataStore(ctrl)

		server := NewPriceListServer(config.ServerConfig{StoreTimeout: time.Second, RestTimeout: time.Hour}, log.NewLogger(), nil, store, run)

		req := &pricelist.ListRequest{
			Limit:            10,
			Offset:           20,
			SortingType:      pricelist.SortingType_SortByProductName,
			SortingDirection: pricelist.SortingDirection_SortAsc,
		}

		prices := []types.ProductPriceExtended{}
		store.EXPECT().Get(gomock.Any(), types.SortByProductName, types.SortAsc, int64(10), int64(20)).Return(prices, nil)

		got, err := server.List(context.Background(), req)
		assert.Nil(t, err)
		assert.Empty(t, got.Products)
	})

	t.Run("when store ok and data not empty", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		store := NewMockMarketDataStore(ctrl)

		server := NewPriceListServer(config.ServerConfig{StoreTimeout: time.Second, RestTimeout: time.Hour}, log.NewLogger(), nil, store, run)

		req := &pricelist.ListRequest{
			Limit:            10,
			Offset:           20,
			SortingType:      pricelist.SortingType_SortByProductName,
			SortingDirection: pricelist.SortingDirection_SortAsc,
		}

		prices := []types.ProductPriceExtended{{
			ProductPrice: types.ProductPrice{
				ProductName:       "AAAA",
				ProductPriceCents: 5678,
				UpdateTimeNano:    time.Unix(1200, 0).UnixNano(),
			},
			UpdatesCount: 12,
		}, {
			ProductPrice: types.ProductPrice{
				ProductName:       "BBB",
				ProductPriceCents: 432,
				UpdateTimeNano:    time.Unix(10, 543678).UnixNano(),
			},
			UpdatesCount: -1,
		}}
		store.EXPECT().Get(gomock.Any(), types.SortByProductName, types.SortAsc, int64(10), int64(20)).Return(prices, nil)

		got, err := server.List(context.Background(), req)
		assert.Nil(t, err)
		assert.Equal(t, []*pricelist.ProductPrices{{
			ProductName:       "AAAA",
			ProductPriceCents: 5678,
			UpdateCount:       12,
			UpdateTime:        "1970-01-01T03:20:00+03:00",
		}, {
			ProductName:       "BBB",
			ProductPriceCents: 432,
			UpdateCount:       -1,
			UpdateTime:        "1970-01-01T03:00:10.000543678+03:00",
		},
		}, got.Products)
	})
}

func TestPriceListServer_Fetch(t *testing.T) {

	t.Run("when fetch data error", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		store := NewMockMarketDataStore(ctrl)
		provider := NewMockMarketDataProvider(ctrl)
		server := NewPriceListServer(config.ServerConfig{StoreTimeout: time.Second, RestTimeout: time.Hour}, log.NewLogger(), provider, store, run)

		req := &pricelist.FetchRequest{Url: "example.com"}

		provider.EXPECT().Fetch(gomock.Any(), "example.com").Return(nil, errors.New("error"))

		got, err := server.Fetch(context.Background(), req)
		assert.NotNil(t, err)
		assert.Nil(t, got)
	})

	t.Run("when fetch data ok", func(t *testing.T) {

		t.Run("when save to store error", func(t *testing.T) {
			// the same as store ok. only log
		})

		t.Run("when save to store ok", func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := NewMockMarketDataStore(ctrl)
			provider := NewMockMarketDataProvider(ctrl)
			server := NewPriceListServer(config.ServerConfig{StoreTimeout: time.Second, RestTimeout: time.Hour}, log.NewLogger(), provider, store, run)

			req := &pricelist.FetchRequest{Url: "example.com"}

			prices := []types.ProductPrice{{}, {}}
			provider.EXPECT().Fetch(gomock.Any(), "example.com").Return(prices, nil)
			store.EXPECT().Save(gomock.Any(), prices).Return(errors.New("error"))

			got, err := server.Fetch(context.Background(), req)
			assert.Nil(t, err)
			assert.Equal(t, []*pricelist.ProductPrice{{}, {}}, got.Products)
		})

	})
}
