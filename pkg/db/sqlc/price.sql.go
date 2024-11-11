// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: price.sql

package sqlc

import (
	"context"
)

const createPrice = `-- name: CreatePrice :one
INSERT INTO price (ad_id, fetched_at, has_price, total_price, price_per_meter, mortgage, normal_price, weekend_price)
VALUES ($1,
        NOW(),
        $2,
        $3,
        $4,
        $5,
        $6,
        $7)
RETURNING id, ad_id, fetched_at, has_price, total_price, price_per_meter, mortgage, normal_price, weekend_price
`

type CreatePriceParams struct {
	AdID          int64  `json:"ad_id"`
	HasPrice      *bool  `json:"has_price"`
	TotalPrice    *int64 `json:"total_price"`
	PricePerMeter *int64 `json:"price_per_meter"`
	Mortgage      *int64 `json:"mortgage"`
	NormalPrice   *int64 `json:"normal_price"`
	WeekendPrice  *int64 `json:"weekend_price"`
}

// Insert a new price entry for a specific ad
func (q *Queries) CreatePrice(ctx context.Context, arg CreatePriceParams) (Price, error) {
	row := q.db.QueryRow(ctx, createPrice,
		arg.AdID,
		arg.HasPrice,
		arg.TotalPrice,
		arg.PricePerMeter,
		arg.Mortgage,
		arg.NormalPrice,
		arg.WeekendPrice,
	)
	var i Price
	err := row.Scan(
		&i.ID,
		&i.AdID,
		&i.FetchedAt,
		&i.HasPrice,
		&i.TotalPrice,
		&i.PricePerMeter,
		&i.Mortgage,
		&i.NormalPrice,
		&i.WeekendPrice,
	)
	return i, err
}
