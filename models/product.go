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

	rows, err := db.SQLDB.Query(`
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

	rows, err := db.SQLDB.Query(querySQL)
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

// ProductPrice 产品价格
type ProductPrice struct {
	ProductID     int     `json:"product_id"`
	ProductClass  string  `json:"product_class"`
	ProductName   string  `json:"product_name"`
	Price         float64 `json:"price"`
	DiscountPrice float64 `json:"discount_price"`
	VipPrice      float64 `json:"vip_price"`
	Period        string  `json:"period"`
	ServiceDay    int     `json:"service_day"`
	ServiceMonth  int     `json:"service_month"`
	Alias         string  `json:"alias"`
	Sort          int     `json:"sort"`
	IsShow        int     `json:"is_show"`
	IsListShow    int     `json:"is_list_show"`
	IsEnable      int     `json:"is_enable"`
	ApplePayPID   string  `json:"apple_pay_pid"`
	IsConsume     int     `json:"is_consume"`
}

// GetProductPrices 查询产品价格
func (p *Product) GetProductPrices() (ppList []ProductPrice, err error) {

	ppList = make([]ProductPrice, 0)

	querySQL := utils.SetSQLFormat(`
	SELECT 
		product_id,product_class,product_name,
		price*0.01,discount_price*0.01,vip_price*0.01,
		period,service_day,service_month,
		IFNULL(alias,""),sort,
		is_show,is_list_show,is_enable,
		IFNULL(apple_pay_pid,""),is_consume
	FROM t_product_price
	WHERE 1
	`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY create_time DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var pp ProductPrice
		rows.Scan(&pp.ProductID, &pp.ProductClass, &pp.ProductName,
			&pp.Price, &pp.DiscountPrice, &pp.VipPrice,
			&pp.Period, &pp.ServiceDay, &pp.ServiceMonth,
			&pp.Alias, &pp.Sort,
			&pp.IsShow, &pp.IsListShow, &pp.IsEnable,
			&pp.ApplePayPID, &pp.IsConsume)

		ppList = append(ppList, pp)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}

// ProductRecommend 推荐产品
type ProductRecommend struct {
	ID              int    `json:"id"`
	RecommendType   string `json:"recommend_type"`
	CNRecommendType string `json:"cn_recommend_type"`
	ProductClass    string `json:"product_class"`
	ProductName     string `json:"product_name"`
	IsList          int    `json:"is_list"`
	Sort            int    `json:"sort"`
}

// GetProductRecommends 查询推荐产品
func (p *Product) GetProductRecommends() (prList []ProductRecommend, err error) {

	prList = make([]ProductRecommend, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT
			prs.id,prs.recommend_type,
			CASE recommend_type
				WHEN 'is_top' THEN '实验室头牌'
				WHEN 'is_new' THEN '新品推荐'
				WHEN 'is_like' THEN '猜你喜欢'
				WHEN 'is_wntj' THEN '为你推荐'
			END AS cn_recommend_type,
			prs.product_class,
			p.product_name,
			prs.is_list,
			prs.sort
		FROM t_product_recommend_setting prs
		LEFT JOIN t_product p ON prs.product_class = p.product_class
		WHERE 1
	`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY prs.id  DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var pr ProductRecommend
		rows.Scan(&pr.ID, &pr.RecommendType, &pr.CNRecommendType,
			&pr.ProductClass, &pr.ProductName,
			&pr.IsList, &pr.Sort)

		prList = append(prList, pr)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}
