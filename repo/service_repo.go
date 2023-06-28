package repo

import (
	"database/sql"
	"fmt"
	"goclean/model"
)

type ServiceRepo interface {
	Get(int) (*model.ServiceModel, error)
	FindByName(string) (*model.ServiceModel, error)
	List() (*[]model.ServiceModel, error)
	Create(*model.ServiceModel) error
}

type serviceRepoImpl struct {
	db *sql.DB
}

func NewServiceRepo(db *sql.DB) ServiceRepo {
	return &serviceRepoImpl{
		db: db,
	}
}

func (svcRepo *serviceRepoImpl) Get(id int) (*model.ServiceModel, error) {
	qry := "SELECT id, name, uom, price FROM ms_service WHERE id = $1"

	svc := &model.ServiceModel{}
	err := svcRepo.db.QueryRow(qry, id).Scan(&svc.Id, &svc.Name, &svc.Uom, &svc.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.Get() : %w", err)
	}
	return svc, nil
}

func (svcRepo *serviceRepoImpl) FindByName(name string) (*model.ServiceModel, error) {
	qry := "SELECT id, name, uom, price FROM ms_service WHERE name = $1"

	svc := &model.ServiceModel{}
	err := svcRepo.db.QueryRow(qry, name).Scan(&svc.Id, &svc.Name, &svc.Uom, &svc.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.GetServiceByName() : %w", err)
	}
	return svc, nil
}

func (svcRepo *serviceRepoImpl) List() (*[]model.ServiceModel, error) {
	qry := "SELECT id, name, uom, price FROM ms_service"
	rows, err := svcRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on serviceRepoImpl.List() : %w", err)
	}
	defer rows.Close()
	var arrSvc []model.ServiceModel
	for rows.Next() {
		svc := &model.ServiceModel{}
		rows.Scan(&svc.Id, &svc.Name, &svc.Price, &svc.Uom)
		arrSvc = append(arrSvc, *svc)
	}
	return &arrSvc, nil
}

func (svcRepo *serviceRepoImpl) Create(scv *model.ServiceModel) error {
	qry := "INSERT INTO ms_service(name, uom, price) VALUES($1, $2, $3)"
	_, err := svcRepo.db.Exec(qry, scv.Name, scv.Uom, scv.Price)
	if err != nil {
		return fmt.Errorf("error on serviceRepoImpl.Create() : %w", err)
	}
	return nil
}
