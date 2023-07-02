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
	qry := "SELECT no, start_date, end_date, cust_name, phone_no FROM tr_header"

	rows, err := trxRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction headers: %v", err)
	}
	defer rows.Close()

	var transactions []model.TransactionHeader

	for rows.Next() {
		header := model.TransactionHeader{}
		err := rows.Scan(&header.No, &header.StartDate, &header.EndDate, &header.CustName, &header.Phone)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction header: %v", err)
		}

		details, err := trxRepo.GetTransactionDetailsByTrxNo(header.No)
		if err != nil {
			return nil, fmt.Errorf("failed to get transaction details for header No %v: %v", header.No, err)
		}
		transaction := model.TransactionHeader{
			No:        header.No,
			StartDate: header.StartDate,
			EndDate:   header.EndDate,
			CustName:  header.CustName,
			Phone:     header.Phone,
			ArrDetail: *details,
		}

		transactions = append(transactions, transaction)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over transaction headers: %v", err)
	}

	return &transactions, nil
}

func (trxRepo *transactionRepoImpl) GetTransactionDetailsByTrxNo(trxNo sql.NullInt64) (*[]model.TransactionDetail, error) {
	qry := "SELECT id,trx_no, service_name, qty, uom, price FROM tr_detail WHERE trx_no = $1"

	rows, err := trxRepo.db.Query(qry, trxNo)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction details: %v", err)
	}
	defer rows.Close()

	var details []model.TransactionDetail

	for rows.Next() {
		det := model.TransactionDetail{}
		err := rows.Scan(&det.Id, &det.No, &det.ServiceName, &det.Qty, &det.Uom, &det.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction detail: %v", err)
		}
		details = append(details, det)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over transaction details: %v", err)
	}

	return &details, nil
}

func (trxRepo *transactionRepoImpl) Create(trx *model.TransactionHeader) error {
	tx, err := trxRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("AddTransaction() Begin : %w", err)
	}

	qry := "INSERT INTO tr_header(start_date, end_date, cust_name, phone_no) VALUES($1, $2, $3, $4) RETURNING no"

	err = tx.QueryRow(qry, trx.StartDate, trx.EndDate, trx.CustName, trx.Phone).Scan(&trx.No)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("AddTransaction() Header : %w", err)
	}

	qry = "INSERT INTO tr_detail(trx_no,service_id,service_name, qty, uom, price) VALUES($1, $2, $3, $4, $5, $6)"
	for _, det := range trx.ArrDetail {
		_, err := tx.Exec(qry, trx.No, det.Id, det.ServiceName, det.Qty, det.Uom, det.Price)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("AddTransaction() Detail : %w", err)
		}
	}

	tx.Commit()

	return nil
}
