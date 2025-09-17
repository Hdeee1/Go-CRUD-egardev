package entities

import "database/sql"

type Product struct {
	Id				uint
	Name			string
	CategoryId 		uint
	Category		Category
	Stock			int64
	Descriptions	string
	Created_at		sql.NullTime
	Updated_at		sql.NullTime
}