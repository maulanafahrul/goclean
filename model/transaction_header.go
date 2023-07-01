package model

import "database/sql"

type TransactionHeader struct {
	No        sql.NullInt64
	StartDate sql.NullTime
	EndDate   sql.NullTime
	CustName  sql.NullString
	Phone     sql.NullString

	ArrDetail []TransactionDetail
}
