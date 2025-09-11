package categorymodel

import (
	"go-crud-egardev/config"
	"go-crud-egardev/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil{
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next(){
		var category entities.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdateAt)

		if err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}