package repositories

import "app/internal/bootstrap/db"

type ScheduleRepositoryImpl struct {
	db *db.Postgres
}

func NewScheduleRepositoryImpl(db *db.Postgres) *ScheduleRepositoryImpl {
	return &ScheduleRepositoryImpl{
		db: db,
	}
}

func (ur *ScheduleRepositoryImpl) GetSmth(token string) {

}
