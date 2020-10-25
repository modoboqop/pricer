package marketdata

import (
	"context"
	"encoding/hex"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"

	"github.com/stretchr/testify/assert"

	"github.com/isavinof/pricer/config"

	"github.com/isavinof/pricer/types"
)

// It's not a unit testing. But faster than integration
func TestMongoStore_Get(t *testing.T) {

	t.SkipNow() // remove this to run tests using db
	ctx := context.Background()
	connection, err := NewMongoConnection(ctx, config.MongoConfig{URL: "mongodb://localhost:27017"})
	assert.Nil(t, err)

	now := time.Now().UnixNano()
	initStore := func() *MongoStore {
		db := connection.Database("db-test")
		uid, _ := uuid.New()
		col := db.Collection("collection-" + hex.EncodeToString(uid[:]))
		col.Drop(ctx)

		store, err := NewMongoStore(ctx, col)
		assert.Nil(t, err)
		return store
	}
	t.Run("sort asc", func(t *testing.T) {
		t.Parallel()

		prices := []types.ProductPrice{{
			ProductName:       "AAA",
			ProductPriceCents: 456,
			UpdateTimeNano:    now,
		}, {
			ProductName:       "BBB",
			ProductPriceCents: 434,
			UpdateTimeNano:    now + int64(time.Second),
		}, {
			ProductName:       "CCC",
			ProductPriceCents: 1,
			UpdateTimeNano:    now - int64(time.Second),
		}}

		store := initStore()
		assert.Nil(t, store.Save(ctx, prices))

		gotPrices, err := store.Get(ctx, types.SortByProductName, types.SortAsc, 4, 0)
		if assert.Nil(t, err) && assert.Len(t, gotPrices, 3) {
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[0], UpdatesCount: 1}, gotPrices[0])
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[1], UpdatesCount: 1}, gotPrices[1])
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[2], UpdatesCount: 1}, gotPrices[2])
		}

		gotPrices, err = store.Get(ctx, types.SortByPrice, types.SortAsc, 4, 0)
		if assert.Nil(t, err) && assert.Len(t, gotPrices, 3) {
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[0], UpdatesCount: 1}, gotPrices[2])
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[1], UpdatesCount: 1}, gotPrices[1])
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[2], UpdatesCount: 1}, gotPrices[0])
		}

		gotPrices, err = store.Get(ctx, types.SortByUpdateTime, types.SortAsc, 4, 0)
		if assert.Nil(t, err) && assert.Len(t, gotPrices, 3) {
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[0], UpdatesCount: 1}, gotPrices[1])
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[1], UpdatesCount: 1}, gotPrices[2])
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[2], UpdatesCount: 1}, gotPrices[0])
		}

	})

	t.Run("sort desc", func(t *testing.T) {
		t.Parallel()

		prices := []types.ProductPrice{{
			ProductName:       "AAA",
			ProductPriceCents: 456,
			UpdateTimeNano:    now,
		}, {
			ProductName:       "BBB",
			ProductPriceCents: 434,
			UpdateTimeNano:    now + int64(time.Second),
		}, {
			ProductName:       "CCC",
			ProductPriceCents: 1,
			UpdateTimeNano:    now - int64(time.Second),
		}}

		store := initStore()
		assert.Nil(t, store.Save(ctx, prices))

		gotPrices, err := store.Get(ctx, types.SortByProductName, types.SortDesc, 4, 0)
		if assert.Nil(t, err) && assert.Len(t, gotPrices, 3) {
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[0], UpdatesCount: 1}, gotPrices[2])
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[1], UpdatesCount: 1}, gotPrices[1])
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[2], UpdatesCount: 1}, gotPrices[0])
		}

		gotPrices, err = store.Get(ctx, types.SortByPrice, types.SortDesc, 4, 0)
		if assert.Nil(t, err) && assert.Len(t, gotPrices, 3) {
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[0], UpdatesCount: 1}, gotPrices[0])
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[1], UpdatesCount: 1}, gotPrices[1])
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[2], UpdatesCount: 1}, gotPrices[2])
		}

		gotPrices, err = store.Get(ctx, types.SortByUpdateTime, types.SortDesc, 4, 0)
		if assert.Nil(t, err) && assert.Len(t, gotPrices, 3) {
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[0], UpdatesCount: 1}, gotPrices[1])
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[1], UpdatesCount: 1}, gotPrices[0])
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: prices[2], UpdatesCount: 1}, gotPrices[2])
		}

	})

	t.Run("sort type multiple updates", func(t *testing.T) {
		t.Parallel()

		db := connection.Database("db-test")
		db.Drop(ctx)
		col := db.Collection("collection-test")
		col.Drop(ctx)

		store, err := NewMongoStore(ctx, col)
		assert.Nil(t, err)

		assert.Nil(t, store.Save(ctx, []types.ProductPrice{{
			ProductName:       "AAA",
			ProductPriceCents: 456,
			UpdateTimeNano:    now,
		}, {
			ProductName:       "BBB",
			ProductPriceCents: 434,
			UpdateTimeNano:    now + int64(time.Second),
		}, {
			ProductName:       "CCC",
			ProductPriceCents: 1,
			UpdateTimeNano:    now - int64(time.Second),
		}}))

		assert.Nil(t, store.Save(ctx, []types.ProductPrice{{
			ProductName:       "AAA",
			ProductPriceCents: 879,
			UpdateTimeNano:    now + int64(time.Hour),
		}, {
			ProductName:       "CCC",
			ProductPriceCents: 195,
			UpdateTimeNano:    now + int64(time.Second+time.Hour),
		}}))

		assert.Nil(t, store.Save(ctx, []types.ProductPrice{{
			ProductName:       "AAA",
			ProductPriceCents: 10789,
			UpdateTimeNano:    now + int64(2*time.Hour),
		}}))

		gotPrices, err := store.Get(ctx, types.SortByUpdatesCount, types.SortAsc, 4, 0)
		if assert.Nil(t, err) && assert.Len(t, gotPrices, 3) {
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: types.ProductPrice{
				ProductName:       "BBB",
				ProductPriceCents: 434,
				UpdateTimeNano:    now + int64(time.Second)}, UpdatesCount: 1}, gotPrices[0])

			assert.Equal(t, types.ProductPriceExtended{ProductPrice: types.ProductPrice{
				ProductName:       "CCC",
				ProductPriceCents: 195,
				UpdateTimeNano:    now + int64(time.Second+time.Hour)}, UpdatesCount: 2}, gotPrices[1])

			assert.Equal(t, types.ProductPriceExtended{ProductPrice: types.ProductPrice{
				ProductName:       "AAA",
				ProductPriceCents: 10789,
				UpdateTimeNano:    now + int64(2*time.Hour)}, UpdatesCount: 3}, gotPrices[2])
		}

		gotPrices, err = store.Get(ctx, types.SortByUpdatesCount, types.SortDesc, 4, 0)
		if assert.Nil(t, err) && assert.Len(t, gotPrices, 3) {
			assert.Equal(t, types.ProductPriceExtended{ProductPrice: types.ProductPrice{
				ProductName:       "BBB",
				ProductPriceCents: 434,
				UpdateTimeNano:    now + int64(time.Second)}, UpdatesCount: 1}, gotPrices[2])

			assert.Equal(t, types.ProductPriceExtended{ProductPrice: types.ProductPrice{
				ProductName:       "CCC",
				ProductPriceCents: 195,
				UpdateTimeNano:    now + int64(time.Second+time.Hour)}, UpdatesCount: 2}, gotPrices[1])

			assert.Equal(t, types.ProductPriceExtended{ProductPrice: types.ProductPrice{
				ProductName:       "AAA",
				ProductPriceCents: 10789,
				UpdateTimeNano:    now + int64(2*time.Hour)}, UpdatesCount: 3}, gotPrices[0])
		}

	})
}
