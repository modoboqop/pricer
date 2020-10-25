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
	ProductNameFieldName  = "_id"
	PriceFieldName        = "price_cents"
	UpdateTimeFieldName   = "update_time_nanosec"
	UpdatesCountFieldName = "updates_count"
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

		{Keys: bson.M{PriceFieldName: 1}, Options: options.Index().SetUnique(false)},
		{Keys: bson.M{PriceFieldName: -1}, Options: options.Index().SetUnique(false)},

		{Keys: bson.M{UpdateTimeFieldName: 1}, Options: options.Index().SetUnique(false)},
		{Keys: bson.M{UpdateTimeFieldName: -1}, Options: options.Index().SetUnique(false)},

		{Keys: bson.M{UpdatesCountFieldName: 1}, Options: options.Index().SetUnique(false)},
		{Keys: bson.M{UpdatesCountFieldName: -1}, Options: options.Index().SetUnique(false)},
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
			SetFilter(bson.M{"_id": price.ProductName}).
			SetUpdate(bson.M{
				"$set": bson.M{
					"price_cents":         price.ProductPriceCents,
					"update_time_nanosec": price.UpdateTimeNano,
				},
				"$inc": bson.M{"updates_count": 1},
			}).
			SetUpsert(true)

		operations = append(operations, operation)
	}

	// TODO: validate bulks result
	_, err := store.collection.BulkWrite(ctx, operations)
	if err != nil {
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
				ProductName:       price[ProductNameFieldName].(string),
				ProductPriceCents: price[PriceFieldName].(int32),
				UpdateTimeNano:    price[UpdateTimeFieldName].(int64),
			},
			UpdatesCount: int64(price[UpdatesCountFieldName].(int32)),
		})
	}

	return prices, nil
}

var sortTypeToFieldName = map[types.SortingType]string{
	types.SortByProductName:  ProductNameFieldName,
	types.SortByPrice:        PriceFieldName,
	types.SortByUpdatesCount: UpdatesCountFieldName,
	types.SortByUpdateTime:   UpdateTimeFieldName,
}

var sortTypeDirectionToInt = map[types.SortDirectionType]int{
	types.SortAsc:  1,
	types.SortDesc: -1,
}
