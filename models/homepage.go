package models

import (
	"log"
	db "restful_gin/database"
	"restful_gin/utils"
)

// Homepage 首页
type Homepage struct{}

// Column 栏目
type Column struct {
	ID             int    `json:"id"`
	ColumnTypeID   int    `json:"column_type_id"`
	ColumnTypeName string `json:"column_type_name"`
	ColumnName     string `json:"column_name"`
	JumpURL        string `json:"jump_url"`
	JumpID         int    `json:"jump_id"`
	Sort           int    `json:"sort"`
	IsShow         int    `json:"is_show"`
	InReview       int    `json:"in_review"`
}

// GetColumns 获取栏目
func (h *Homepage) GetColumns() (columns []Column, err error) {
	columns = make([]Column, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id,
			column_type_id,column_type_name,
			column_name ,jump_url,jump_id,
			sort,is_show,in_review
		FROM t_column 
		WHERE 1
	`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}
	for rows.Next() {
		var c Column
		rows.Scan(&c.ID, &c.ColumnTypeID, &c.ColumnTypeName,
			&c.ColumnName, &c.JumpURL, &c.JumpID,
			&c.Sort, &c.IsShow, &c.InReview)

		columns = append(columns, c)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}

// ProductColumn 首页产品推荐栏目
type ProductColumn struct {
	ID             int     `json:"id"`
	ColumnID       int     `json:"column_id"`
	ColumnName     string  `json:"column_name"`
	ColumnTypeName string  `json:"column_type_name"`
	ImageURL       string  `json:"image_url"`
	ProductName    string  `json:"product_name"`
	ProductClass   string  `json:"product_class"`
	ProductDesc    string  `json:"product_desc"`
	SubscribeNum   int     `json:"subscribe_num"`
	Price          float64 `json:"price"`
	IsShow         int     `json:"is_show"`
}

// GetProductColumns 查询首页产品推荐栏目
func (h *Homepage) GetProductColumns() (pcs []ProductColumn, err error) {

	pcs = make([]ProductColumn, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT  hpc.id,
			hpc.column_id,
			c.column_name,
			c.column_type_name,
			hpc.image_url ,
			hpc.product_name,
			hpc.product_class,
			hpc.product_desc ,
			hpc.subscribe_num ,
			hpc.price,
			hpc.is_show
		FROM t_homepage_product_recommend hpc
		LEFT JOIN  t_column c ON c.id = hpc.column_id
		WHERE 1
	`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}
	for rows.Next() {
		var pc ProductColumn
		rows.Scan(&pc.ID, &pc.ColumnID, &pc.ColumnName, &pc.ColumnTypeName,
			&pc.ImageURL, &pc.ProductName, &pc.ProductClass, &pc.ProductDesc,
			&pc.SubscribeNum, &pc.Price, &pc.IsShow)

		pcs = append(pcs, pc)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}

// SPColumn 荐股产品栏目
type SPColumn struct {
	ID                 int    `json:"id"`
	ColumnID           int    `json:"column_id"`
	ColumnName         string `json:"column_name"`
	ColumnTypeName     string `json:"column_type_name"`
	StockName          string `json:"stock_name"`
	StockCode          string `json:"stock_code"`
	ProductName        string `json:"product_name"`
	ProductClass       string `json:"product_class"`
	JumpURL            string `json:"jump_url"`
	StockTrendImageURL string `json:"stock_trend_image_url"`
	IsShow             int    `json:"is_show"`
}

// GetSPColumns 查询 首页荐股产品栏目
func (h *Homepage) GetSPColumns() (spcs []SPColumn, err error) {

	spcs = make([]SPColumn, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT 
			hpsr.id,
			hpsr.column_id,
			c.column_name,
			c.column_type_name,
			hpsr.stock_name,
			hpsr.stock_code,
			hpsr.product_name,
			hpsr.product_class,
			hpsr.jump_url,
			hpsr.stock_trend_image_url,
			hpsr.is_show
		FROM gpxj_app.t_homepage_product_stock_recommend hpsr
		LEFT JOIN gpxj_app.t_column c ON hpsr.column_id = c.id
		WHERE 1
	`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}
	for rows.Next() {
		var spc SPColumn
		rows.Scan(&spc.ID, &spc.ColumnID, &spc.ColumnName, &spc.ColumnTypeName,
			&spc.StockName, &spc.StockCode, &spc.ProductName, &spc.ProductClass,
			&spc.JumpURL, &spc.StockTrendImageURL, &spc.IsShow)

		spcs = append(spcs, spc)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}

// ArticleColumn 文章栏目
type ArticleColumn struct {
	ID             int    `json:"id"`
	ColumnID       int    `json:"column_id"`
	ColumnName     string `json:"column_name"`
	ColumnTypeName string `json:"column_type_name"`
	IconURL        string `json:"icon_url"`
	Nickname       string `json:"nickname"`
	JumpURL        string `json:"jump_url"`
	IsShow         int    `json:"is_show"`
}

// GetArticleColumns 查询首页文章栏目
func (h *Homepage) GetArticleColumns() (acs []ArticleColumn, err error) {

	acs = make([]ArticleColumn, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT 
			ha.id,
			ha.column_id ,
			c.column_name,
			c.column_type_name,
			ha.icon_url ,
			ha.nickname,
			ha.jump_url,
			ha.is_show
		FROM t_homepage_article ha
		LEFT JOIN t_column c ON ha.column_id = c.id
		WHERE 1
	`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}
	for rows.Next() {
		var ac ArticleColumn
		rows.Scan(&ac.ID, &ac.ColumnID, &ac.ColumnName, &ac.ColumnTypeName,
			&ac.IconURL, &ac.Nickname, &ac.JumpURL, &ac.IsShow)

		acs = append(acs, ac)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}

// ShortCutMenu 快捷菜单
type ShortCutMenu struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	IconURL      string `json:"icon_url"`
	JumpURL      string `json:"jump_url"`
	IsNew        int    `json:"is_new"`
	InReview     int    `json:"in_review"`
	ProductClass string `json:"product_class"`
}

// GetShortCutMenus 查询首页快捷菜单
func (h *Homepage) GetShortCutMenus() (scms []ShortCutMenu, err error) {

	scms = make([]ShortCutMenu, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id, title,
			icon_url, jump_url,
			is_new, in_review, product_class  
		FROM gpxj_app.t_homepage_product_menu
		WHERE 1
		`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var scm ShortCutMenu
		rows.Scan(&scm.ID, &scm.Title, &scm.IconURL, &scm.JumpURL, &scm.IsNew, &scm.InReview, &scm.ProductClass)

		scms = append(scms, scm)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}

// ProductClassify 产品分类
type ProductClassify struct {
	ID              int    `json:"id"`
	ProductTypeName string `json:"product_type_name"`
	ProductTypeDesc string `json:"product_type_desc"`
	ProductClass    string `json:"product_class"`
	SmallIconURL    string `json:"small_icon_url"`
	ListIconURL     string `json:"list_icon_url"`
	InnerJumpURL    string `json:"inner_jump_url"`
	HasNew          int    `json:"has_new"`
	InReview        int    `json:"in_review"`
	IsShow          int    `json:"is_show"`
}

// GetProductClassifys 查询首页产品分类
func (h *Homepage) GetProductClassifys() (pcs []ProductClassify, err error) {
	pcs = make([]ProductClassify, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id,
			product_type_name, product_type_desc, product_class,
			small_icon_url, list_icon_url, inner_jump_url,
			has_new,in_review, is_show 
		FROM gpxj_app.t_product_type 
		WHERE 1
	`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}
	for rows.Next() {
		var pc ProductClassify
		rows.Scan(&pc.ID, &pc.ProductTypeName, &pc.ProductTypeDesc, &pc.ProductClass,
			&pc.SmallIconURL, &pc.ListIconURL, &pc.InnerJumpURL,
			&pc.HasNew, &pc.InReview, &pc.IsShow)

		pcs = append(pcs, pc)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}

// ActivityMarketing 活动营销
type ActivityMarketing struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	ImageURL     string `json:"image_url"`
	IconURL      string `json:"icon_url"`
	JumpURL      string `json:"jump_url"`
	Uptime       string `json:"uptime"`
	ActivityType int    `json:"activity_type"`
	IsShow       int    `json:"is_show"`
}

// GetActivityMarketings 查询活动营销
func (h *Homepage) GetActivityMarketings(mType int) (ams []ActivityMarketing, err error) {
	ams = make([]ActivityMarketing, 0)
	querySQL := utils.SetSQLFormat(`
		SELECT id,title,
			IFNULL(image_url,""),
			IFNULL(icon_url,""), 
			IFNULL(jump_url,""),
			uptime,activity_type,is_show
		FROM t_activity_marketing
		WHERE 1
	`)

	if 1 == mType {
		querySQL = utils.SetSQLFormat(`{0} AND marketing_type = 'suspension'`, querySQL)
	} else if 2 == mType {
		querySQL = utils.SetSQLFormat(`{0} AND marketing_type = 'popups'`, querySQL)
	}

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var am ActivityMarketing
		rows.Scan(&am.ID, &am.Title, &am.ImageURL, &am.IconURL, &am.JumpURL,
			&am.Uptime, &am.ActivityType, &am.IsShow)

		ams = append(ams, am)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}
