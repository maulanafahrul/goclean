package model

import "database/sql"

type TransactionDetail struct {
	Id          sql.NullInt64
	No          sql.NullInt64
	ServiceId   sql.NullInt64
	ServiceName sql.NullString
	Qty         sql.NullFloat64
	Price       sql.NullFloat64
	Uom         sql.NullString
}
