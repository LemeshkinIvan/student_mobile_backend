package repositories

import "app/internal/bootstrap/db"

type NewsRepositoryImpl struct {
	db *db.Postgres
}

func NewNewsRepository(db *db.Postgres) *NewsRepositoryImpl {
	return &NewsRepositoryImpl{
		db: db,
	}
}

func (ur *NewsRepositoryImpl) GetSmth(token string) {

}
