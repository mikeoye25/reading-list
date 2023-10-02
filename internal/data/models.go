package data

import "database/sql"

type Models struct {
	Books BookModel
	Users UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Books: BookModel{DB: db},
		Users: UserModel{DB: db},
	}
}
