package models

import db "restful_gin/database"

// Product 产品
type Product struct {
	ProductClass string `json:"product_class"`
	ProductName  string `json:"product_name"`
}

// GetProducts 单独产品
func (p *Product) GetProducts() (products []Product, err error) {
	products = make([]Product, 0)

	rows, err := db.SqlDB.Query(`
		SELECT product_class,product_name
		FROM gpxj_app.t_product
		WHERE is_package <> 1
	`)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var product Product
		rows.Scan(&product.ProductClass, &product.ProductName)

		products = append(products, product)
	}
	if err = rows.Err(); nil != err {
		return
	}
	return
}
