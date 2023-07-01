package repo

import (
	"database/sql"
	"fmt"
	"goclean/model"
)

type TransactionRepo interface {
	List() (*[]model.TransactionHeader, error)
	Create(*model.TransactionHeader) error
}

type transactionRepoImpl struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) TransactionRepo {
	return &transactionRepoImpl{
		db: db,
	}
}

func (trxRepo *transactionRepoImpl) List() (*[]model.TransactionHeader, error) {
	qry := "SELECT no, start_date, end_date, cust_name, phone_no ,id ,trx_no ,service_id ,service_name, qty, uom, price FROM tr_header LEFT JOIN tr_detail ON no = trx_no ORDER BY no"

	rows, err := trxRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on transactionRepoImpl.List() : %w", err)
	}
	defer rows.Close()

	var arrayTransaction []model.TransactionHeader
	for rows.Next() {
		trx := model.TransactionHeader{}
		detail := model.TransactionDetail{}
		err := rows.Scan(
			&trx.No, &trx.StartDate, &trx.EndDate, &trx.CustName, &trx.Phone, &detail.Id, &detail.No, &detail.ServiceId.Id,
			&detail.ServiceId.Name, &detail.Qty, &detail.ServiceId.Uom, &detail.ServiceId.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("error on transactionRepoImpl.List(): %w", err)
		}
		trx.ArrDetail = append(trx.ArrDetail, detail)
		arrayTransaction = append(arrayTransaction, trx)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error on transactionRepoImpl.List(): %w", err)
	}
	return &arrayTransaction, nil
}

func (trxRepo *transactionRepoImpl) Create(trx *model.TransactionHeader) error {
	tx, err := trxRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("error on transactionRepoImpl.Create() Begin : %w", err)
	}

	qry := "INSERT INTO tr_header(start_date, end_date, cust_name, phone_no) VALUES($1, $2, $3, $4) RETURNING no"

	err = tx.QueryRow(qry, trx.StartDate, trx.EndDate, trx.CustName, trx.Phone).Scan(&trx.No)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error on transactionRepoImpl.Create() Header : %w", err)
	}

	qry = "INSERT INTO tr_detail(trx_no, service_id, service_name, qty, uom, price) VALUES($1, $2, $3, $4, $5)"
	for _, det := range trx.ArrDetail {
		_, err := tx.Exec(qry, trx.No, det.ServiceId, det.ServiceId.Name, det.Qty, det.ServiceId.Uom, det.ServiceId.Price)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error on transactionRepoImpl.Create() Detail : %w", err)
		}
	}

	tx.Commit()

	return nil
}
