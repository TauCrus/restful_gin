package models

import (
	"log"
	db "restful_gin/database"
	"restful_gin/utils"
)

// Product 产品
type Product struct {
	ProductClass string `json:"product_class"`
	ProductName  string `json:"product_name"`
}

// GetProducts 单品
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

// ProductList 产品列表参数
type ProductList struct {
	ID                   string `json:"id"`
	ProductClass         string `json:"product_class"`
	ProductName          string `json:"product_name"`
	ProductDesc          string `json:"product_desc"`
	ProductIconURL       string `json:"product_icon_url"`
	ProductJumpURL       string `json:"product_jump_url"`
	ProductImageURL      string `json:"product_image_url"`
	ProductLargeImageURL string `json:"product_large_image_url"`
	ProductWebURL        string `json:"product_web_url"`
	BaseSubscription     int    `json:"base_subscription"`
	Version              string `json:"version"`
	Uptime               string `json:"uptime"`
	IsPackage            int    `json:"is_package"`
	IsNew                int    `json:"is_new"`
	IsShow               int    `json:"is_show"`
	IsJump               int    `json:"is_jump"`
	IsWebJump            int    `json:"is_web_jump"`
	IsApplePay           int    `json:"is_apple_pay"`
	IsSHowInReview       int    `json:"is_show_in_review"`
	IsPresale            int    `json:"is_presale"`
	Sort                 int    `json:"sort"`
	ProductTypeID        int    `json:"product_type_id"`
	ProductTypeName      string `json:"product_type_name"`
}

// GetProductList 产品列表
func (p *Product) GetProductList() (productList []ProductList, err error) {
	productList = make([]ProductList, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT  p.id,
		p.product_class,
		p.product_name,
		p.product_desc,
		IFNULL(p.product_icon_url,""),
		IFNULL(p.product_jump_url,""),
		IFNULL(p.product_image_url,""),
		IFNULL(p.product_large_image_url,""),
		IFNULL(p.product_web_url,""),
		IFNULL(p.base_subscription,""),
		IFNULL(p.version,""),
		IFNULL(p.uptime,""),
		p.is_package,
		p.is_new,
		p.is_show,
		p.is_jump,
		p.is_web_jump,
		p.is_apple_pay,
		p.is_show_in_review,
		p.is_presale,
		p.sort,
		p.product_type_id,
		IFNULL(pt.product_type_name,"")
		FROM t_product p
		LEFT JOIN t_product_type pt ON  p.product_type_id = pt.id 
	`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY p.create_time DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SqlDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var pl ProductList
		rows.Scan(&pl.ID, &pl.ProductClass, &pl.ProductName, &pl.ProductDesc,
			&pl.ProductIconURL, &pl.ProductJumpURL, &pl.ProductImageURL, &pl.ProductLargeImageURL, &pl.ProductWebURL,
			&pl.BaseSubscription, &pl.Version, &pl.Uptime,
			&pl.IsPackage, &pl.IsNew, &pl.IsShow, &pl.IsJump, &pl.IsWebJump, &pl.IsApplePay, &pl.IsSHowInReview, &pl.IsPresale,
			&pl.Sort, &pl.ProductTypeID, &pl.ProductTypeName)

		productList = append(productList, pl)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}
