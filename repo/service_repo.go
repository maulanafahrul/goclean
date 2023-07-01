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
	Update(*model.ServiceModel) error
	Delete(int) error
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
		rows.Scan(&svc.Id, &svc.Name, &svc.Uom, &svc.Price)
		arrSvc = append(arrSvc, *svc)
	}
	return &arrSvc, nil
}

func (svcRepo *serviceRepoImpl) Create(svc *model.ServiceModel) error {
	qry := "INSERT INTO ms_service(name, uom, price) VALUES($1, $2, $3)"
	_, err := svcRepo.db.Exec(qry, svc.Name, svc.Uom, svc.Price)
	if err != nil {
		return fmt.Errorf("error on serviceRepoImpl.Create() : %w", err)
	}
	return nil
}

func (svcRepo *serviceRepoImpl) Update(svc *model.ServiceModel) error {
	qry := "UPDATE ms_service SET name=$1, uom=$2,  price=$3 WHERE id=$4"
	_, err := svcRepo.db.Exec(qry, svc.Name, svc.Uom, svc.Price, svc.Id)
	if err != nil {
		return fmt.Errorf("error on serviceRepoImpl.Update() : %w", err)
	}
	return nil
}

func (svcRepo *serviceRepoImpl) Delete(id int) error {
	qry := "DELETE FROM ms_service WHERE id = $1"
	_, err := svcRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("error on serviceRepoImpl.Delete() : %w", err)
	}
	return nil
}
