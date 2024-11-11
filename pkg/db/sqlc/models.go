// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type AdCategory string

const (
	AdCategoryRent     AdCategory = "rent"
	AdCategoryBuy      AdCategory = "buy"
	AdCategoryMortgage AdCategory = "mortgage"
	AdCategoryOther    AdCategory = "other"
)

func (e *AdCategory) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AdCategory(s)
	case string:
		*e = AdCategory(s)
	default:
		return fmt.Errorf("unsupported scan type for AdCategory: %T", src)
	}
	return nil
}

type NullAdCategory struct {
	AdCategory AdCategory
	Valid      bool // Valid is true if AdCategory is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAdCategory) Scan(value interface{}) error {
	if value == nil {
		ns.AdCategory, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AdCategory.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAdCategory) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AdCategory), nil
}

type HouseType string

const (
	HouseTypeApartment HouseType = "apartment"
	HouseTypeVilla     HouseType = "villa"
	HouseTypeOther     HouseType = "other"
)

func (e *HouseType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = HouseType(s)
	case string:
		*e = HouseType(s)
	default:
		return fmt.Errorf("unsupported scan type for HouseType: %T", src)
	}
	return nil
}

type NullHouseType struct {
	HouseType HouseType
	Valid     bool // Valid is true if HouseType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullHouseType) Scan(value interface{}) error {
	if value == nil {
		ns.HouseType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.HouseType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullHouseType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.HouseType), nil
}

type UserRole string

const (
	UserRoleSuperAdmin UserRole = "super_admin"
	UserRoleAdmin      UserRole = "admin"
	UserRoleSimple     UserRole = "simple"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole
	Valid    bool // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Ad struct {
	ID             int64
	PublisherAdKey string
	PublisherID    pgtype.Int4
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
	PublishedAt    pgtype.Timestamp
	Category       NullAdCategory
	Author         pgtype.Text
	Url            pgtype.Text
	Title          pgtype.Text
	Description    pgtype.Text
	City           pgtype.Text
	Neighborhood   pgtype.Text
	HouseType      NullHouseType
	Meterage       pgtype.Int4
	RoomsCount     pgtype.Int4
	Year           pgtype.Int4
	Floor          pgtype.Int4
	TotalFloors    pgtype.Int4
	HasWarehouse   pgtype.Bool
	HasElevator    pgtype.Bool
	Lat            pgtype.Numeric
	Lng            pgtype.Numeric
}

type AdPicture struct {
	ID   int64
	AdID pgtype.Int8
	Url  pgtype.Text
}

type FavoriteAd struct {
	ID     int64
	UserID pgtype.Text
	AdID   pgtype.Int8
}

type Price struct {
	ID            int32
	AdID          pgtype.Int8
	FetchedAt     pgtype.Timestamp
	HasPrice      pgtype.Bool
	TotalPrice    pgtype.Int8
	PricePerMeter pgtype.Int8
	Mortgage      pgtype.Int8
	NormalPrice   pgtype.Int8
	WeekendPrice  pgtype.Int8
}

type Publisher struct {
	ID   int32
	Name string
	Url  string
}

type User struct {
	TgID            string
	Role            NullUserRole
	WatchlistPeriod pgtype.Int4
}

type UserAd struct {
	UserID pgtype.Text
	AdID   pgtype.Int8
}
