package repo

import "database/sql"

type ServiceRepo interface {
}

type serviceRepoImpl struct {
	db *sql.DB
}

func NewServiceRepo(db *sql.DB) ServiceRepo {
	return &serviceRepoImpl{
		db: db,
	}
}
