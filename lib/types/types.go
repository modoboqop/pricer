package types

// ProductPrice without updates count
type ProductPrice struct {
	ProductName       string
	ProductPriceCents int32
	UpdateTimeNano    int64
}

// ProductPrice with updates count
type ProductPriceExtended struct {
	ProductPrice
	UpdatesCount int64
}

// SortingType the same as in grpc server
type SortingType string

const (
	SortByProductName  SortingType = "SortByProductName"
	SortByPrice        SortingType = "SortByPrice"
	SortByUpdatesCount SortingType = "SortByUpdatesCount"
	SortByUpdateTime   SortingType = "SortByUpdateTime"
)

// SortDirectionType direction of sorting
type SortDirectionType string

const (
	SortAsc  SortDirectionType = "SortAsc"
	SortDesc SortDirectionType = "SortDesc"
)
