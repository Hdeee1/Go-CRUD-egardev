package productmodel

import (
	"go-crud-egardev/config"
	"go-crud-egardev/entities"
)

func GetAll() []entities.Product{
	rows, err := config.DB.Query(`
	SELECT 
		products.id,
		categories.id as category_name,
		products.stock,
		products.descriptions,
		products.created_at,
		products.updated_at
	FROM products
	JOIN categories ON products.category_id = category.id
    `)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next(){
		var product entities.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Name,
			&product.Stock,
			&product.Descriptions,
			&product.Created_at,
			&product.Updated_at,
		)

		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}