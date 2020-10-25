package marketdata

import (
	"context"
	"fmt"

	"github.com/isavinof/pricer/config"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/isavinof/pricer/types"
)

// MongoStore conform store interface
// allow save and request price data
type MongoStore struct {
	collection *mongo.Collection
}

const (
	productNameFieldName  = "_id"
	priceFieldName        = "price_cents"
	updateTimeFieldName   = "update_time_nanosec"
	updatesCountFieldName = "updates_count"

	duplicateKeyErrorCode = 11000
)

// NewMongoConnection
func NewMongoConnection(ctx context.Context, mongoConfig config.MongoConfig) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoConfig.URL)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, errors.Wrapf(err, "connect to db. url:%v", mongoConfig.URL)
	}

	return client, nil
}

// NewMongoStore create needed indexes
func NewMongoStore(ctx context.Context, collection *mongo.Collection) (*MongoStore, error) {
	if collection == nil {
		return nil, errors.New("collection must be initialized")
	}
	_, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{

		{Keys: bson.M{priceFieldName: 1}, Options: options.Index().SetUnique(false)},
		{Keys: bson.M{priceFieldName: -1}, Options: options.Index().SetUnique(false)},

		{Keys: bson.M{updateTimeFieldName: 1}, Options: options.Index().SetUnique(false)},
		{Keys: bson.M{updateTimeFieldName: -1}, Options: options.Index().SetUnique(false)},

		{Keys: bson.M{updatesCountFieldName: 1}, Options: options.Index().SetUnique(false)},
		{Keys: bson.M{updatesCountFieldName: -1}, Options: options.Index().SetUnique(false)},
	})
	if err != nil {
		return nil, errors.Wrap(err, "create indexes")
	}

	return &MongoStore{collection: collection}, nil
}

// Save storing price data to mongo
func (store *MongoStore) Save(ctx context.Context, prices []types.ProductPrice) error {
	operations := make([]mongo.WriteModel, 0, len(prices))
	for _, price := range prices {
		operation := mongo.NewUpdateOneModel().
			// update product only if we don't have higher time and equal price
			SetFilter(bson.M{
				productNameFieldName: price.ProductName,
				updateTimeFieldName:  bson.M{"$lt": price.UpdateTimeNano},
				priceFieldName:       bson.M{"$ne": price.ProductPriceCents},
			}).
			SetUpdate(bson.M{
				"$set": bson.M{
					priceFieldName:      price.ProductPriceCents,
					updateTimeFieldName: price.UpdateTimeNano,
				},
				"$inc": bson.M{updatesCountFieldName: 1},
			}).
			SetUpsert(true)

		operations = append(operations, operation)
	}

	_, err := store.collection.BulkWrite(ctx, operations, (&options.BulkWriteOptions{}).SetOrdered(false))
	if err != nil {
		// Duplicate key error is ok in this flow
		// on upsert price searching by time and price as well.
		// mongo will try insert value only in 2 ways
		// 1. Object with _id not exists: will be inserted without error
		// 2. Object with _id exists but don't fit to the price and time condition.
		// Second situation will be errored with Duplicate Key error and should be ignored
		if blkErr, ok := err.(mongo.BulkWriteException); ok {
			for _, err := range blkErr.WriteErrors {
				if err.Code != duplicateKeyErrorCode {
					return errors.Wrap(err, "save to db")
				}
			}
			return nil
		}
		return errors.Wrap(err, "save to db")
	}

	return nil
}

// Get get data directly from mongo without caches
func (store *MongoStore) Get(ctx context.Context, sortType types.SortingType, direction types.SortDirectionType, limit int64, offset int64) (prices []types.ProductPriceExtended, err error) {
	opts := options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
		Sort:  bson.M{sortTypeToFieldName[sortType]: sortTypeDirectionToInt[direction]},
	}

	fmt.Println("opts:", opts)

	cursor, err := store.collection.Find(ctx, bson.D{}, &opts)
	if err != nil {
		return nil, errors.Wrap(err, "get from DB")
	}

	defer cursor.Close(ctx)

	var dbPrices = []bson.M{}
	err = cursor.All(ctx, &dbPrices)
	if err != nil {
		return nil, errors.Wrap(err, "read results from DB")
	}

	prices = make([]types.ProductPriceExtended, 0, len(dbPrices))
	for _, price := range dbPrices {
		// TODO: each field validation

		prices = append(prices, types.ProductPriceExtended{
			ProductPrice: types.ProductPrice{
				ProductName:       price[productNameFieldName].(string),
				ProductPriceCents: price[priceFieldName].(int32),
				UpdateTimeNano:    price[updateTimeFieldName].(int64),
			},
			UpdatesCount: int64(price[updatesCountFieldName].(int32)),
		})
	}

	return prices, nil
}

var sortTypeToFieldName = map[types.SortingType]string{
	types.SortByProductName:  productNameFieldName,
	types.SortByPrice:        priceFieldName,
	types.SortByUpdatesCount: updatesCountFieldName,
	types.SortByUpdateTime:   updateTimeFieldName,
}

var sortTypeDirectionToInt = map[types.SortDirectionType]int{
	types.SortAsc:  1,
	types.SortDesc: -1,
}
