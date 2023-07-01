package model

import "database/sql"

type TransactionDetail struct {
	Id        sql.NullInt64
	No        sql.NullInt64
	ServiceId ServiceModel
	Qty       sql.NullFloat64
}
