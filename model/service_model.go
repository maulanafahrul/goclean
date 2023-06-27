package model

import "database/sql"

type ServiceModel struct {
	Id    sql.NullInt64
	Name  sql.NullString
	Uom   sql.NullString
	Price sql.NullFloat64
}
