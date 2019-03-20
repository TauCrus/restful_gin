package models

import (
	"log"
	db "restful_gin/database"
	"restful_gin/utils"

	"github.com/golang/glog"
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

// AddColumn 新增栏目
func (h *Homepage) AddColumn(column Column) (id int64, err error) {

	insertSQL := utils.SetSQLFormat(`
		INSERT INTO gpxj_app.t_column (
			column_type_id,column_type_name,column_name,
			jump_url,jump_id,
			sort,is_show,in_review) 
		VALUES
			(
			'{0}','{1}','{2}',
			'{3}','{4}',
			'{5}','{6}','{7}'
			); 
		`, column.ColumnTypeID, column.ColumnTypeName, column.ColumnName,
		column.JumpURL, column.JumpID,
		column.Sort, column.IsShow, column.InReview)

	glog.Info("insertSQL:", insertSQL)

	rs, err := db.SQLDB.Exec(insertSQL)
	if nil != err {
		return
	}

	id, err = rs.LastInsertId()
	if nil != err {
		return
	}

	return
}

// ModifyColumn 修改栏目
func (h *Homepage) ModifyColumn(column Column) (id int64, err error) {
	updateSQL := utils.SetSQLFormat(`
		UPDATE 
			gpxj_app.t_column 
		SET
			column_type_id = '{1}',
			column_type_name = '{2}',
			column_name = '{3}',
			jump_url = '{4}',
			jump_id = '{5}',
			sort = '{6}',
			is_show = '{7}',
			in_review = '{8}'
		WHERE id = '{0' ;
		`, column.ID,
		column.ColumnTypeID, column.ColumnTypeName, column.ColumnName,
		column.JumpURL, column.JumpID,
		column.Sort, column.IsShow, column.InReview)

	glog.Info("updateSQL:", updateSQL)

	stmt, err := db.SQLDB.Prepare(updateSQL)
	if nil != err {
		return
	}

	rs, err := stmt.Exec()
	if nil != err {
		return
	}

	id, err = rs.RowsAffected()
	if nil != err {
		return
	}

	return
}

// DropColumn 删除栏目
func (h *Homepage) DropColumn(column Column) (id int64, err error) {

	deleteSQL := utils.SetSQLFormat(`DELETE FROM gpxj_app.t_column WHERE id ='{0}'`, column.ID)
	glog.Info("deleteSQL:", deleteSQL)

	rs, err := db.SQLDB.Exec(deleteSQL)
	if nil != err {
		return
	}
	id, err = rs.RowsAffected()
	if nil != err {
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
	JumpURL        string  `json:"jump_url"`
	Sort           int     `json:"sort"`
	IsShow         int     `json:"is_show"`

	ShareTitle   string `json:"share_title"`
	ShareIconURL string `json:"share_icon_url"`
	ShareURL     string `json:"share_url"`
	ShareDesc    string `json:"share_desc"`
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

// AddProductColumn 新增产品推荐栏目
func (h *Homepage) AddProductColumn(pc ProductColumn) (id int64, err error) {

	insertSQL := utils.SetSQLFormat(`
	 INSERT INTO gpxj_app.t_homepage_product_recommend (
		column_id,image_url,
		product_name,product_class,product_desc,
		subscribe_num,price,jump_url,
		share_title,share_desc,share_icon_url,share_url,
		sort,is_show
	  ) 
	  VALUES
		(
		  '{0}','{1}',
		  '{2}','{3}','{4}',
		  '{5}','{6}','{7}',
		  '{8}','{9}','{10}','{11}',
		  '{12}','{13}'
		);
		`, pc.ColumnID, pc.ImageURL,
		pc.ProductName, pc.ProductClass, pc.ProductDesc,
		pc.SubscribeNum, pc.Price, pc.JumpURL,
		pc.ShareTitle, pc.ShareDesc, pc.ShareIconURL, pc.ShareURL,
		pc.Sort, pc.IsShow)

	glog.Info("insertSQL:", insertSQL)

	rs, err := db.SQLDB.Exec(insertSQL)
	if nil != err {
		return
	}

	id, err = rs.LastInsertId()
	if nil != err {
		return
	}

	return
}

// ModifyProductColumn 修改产品推荐栏目
func (h *Homepage) ModifyProductColumn(pc ProductColumn) (id int64, err error) {
	updateSQL := utils.SetSQLFormat(`
	UPDATE 
		gpxj_app.t_homepage_product_recommend 
	SET
		column_id = '{1}',
		image_url = '{2}',
		product_name = '{3}',
		product_class = '{4}',
		product_desc = '{5}',
		subscribe_num = '{6}',
		price = '{7}',
		jump_url = '{8}',
		share_title = '{9}',
		share_desc = '{10}',
		share_icon_url = '{11}',
		share_url = '{12}',
		sort = '{13}',
		is_show = '{14}'
	WHERE id = '{0}' ;
	`, pc.ID,
		pc.ColumnID, pc.ImageURL,
		pc.ProductName, pc.ProductClass, pc.ProductDesc,
		pc.SubscribeNum, pc.Price, pc.JumpURL,
		pc.ShareTitle, pc.ShareDesc, pc.ShareIconURL, pc.ShareURL,
		pc.Sort, pc.IsShow)

	glog.Info("updateSQL:", updateSQL)

	stmt, err := db.SQLDB.Prepare(updateSQL)
	if nil != err {
		return
	}

	rs, err := stmt.Exec()
	if nil != err {
		return
	}

	id, err = rs.RowsAffected()
	if nil != err {
		return
	}

	return
}

// DropProductColumn 删除产品推荐栏目
func (h *Homepage) DropProductColumn(pc ProductColumn) (id int64, err error) {

	deleteSQL := utils.SetSQLFormat(`DELETE FROM gpxj_app.t_homepage_product_recommend WHERE id ='{0}'`, pc.ID)
	glog.Info("deleteSQL:", deleteSQL)

	rs, err := db.SQLDB.Exec(deleteSQL)
	if nil != err {
		return
	}
	id, err = rs.RowsAffected()
	if nil != err {
		return
	}
	return
}

// SPColumn 荐股产品栏目
type SPColumn struct {
	ID                 int     `json:"id"`
	ColumnID           int     `json:"column_id"`
	ColumnName         string  `json:"column_name"`
	ColumnTypeName     string  `json:"column_type_name"`
	StockName          string  `json:"stock_name"`
	StockCode          string  `json:"stock_code"`
	SelectPrice        float64 `json:"select_price"`
	SelectTime         string  `json:"select_time"`
	ProfitDesc         string  `json:"profit_desc"`
	ProfitRatio        float64 `json:"profit_ratio"`
	ProductName        string  `json:"product_name"`
	ProductClass       string  `json:"product_class"`
	ProductDesc        string  `json:"product_desc"`
	JumpURL            string  `json:"jump_url"`
	StockTrendImageURL string  `json:"stock_trend_image_url"`
	Sort               int     `json:"sort"`
	IsShow             int     `json:"is_show"`

	ShareTitle   string `json:"share_title"`
	ShareIconURL string `json:"share_icon_url"`
	ShareURL     string `json:"share_url"`
	ShareDesc    string `json:"share_desc"`
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

// AddSPColumn 新增栏目
func (h *Homepage) AddSPColumn(spc SPColumn) (id int64, err error) {

	insertSQL := utils.SetSQLFormat(`
	INSERT INTO gpxj_app.t_homepage_product_stock_recommend (
		column_id,stock_name,stock_code,
		select_price,select_time,profit_desc,profit_ratio,
		product_name,product_class,product_desc,
		jump_url,stock_trend_image_url,
		share_title,share_desc,share_icon_url,share_url,
		sort,is_show) 
	  VALUES
		(
		  '{0}','{1}','{2}',
		  '{3}','{4}','{5}','{6}',
		  '{7}','{8}','{9}',
		  '{10}','{11}',
		  '{12}','{13}','{14}','{15}',
		  '{16}','{17}');
	`, spc.ColumnID, spc.StockName, spc.StockCode,
		spc.SelectPrice, spc.SelectTime, spc.ProfitDesc, spc.ProfitRatio,
		spc.ProductName, spc.ProductClass, spc.ProductDesc,
		spc.JumpURL, spc.StockTrendImageURL,
		spc.ShareTitle, spc.ShareDesc, spc.ShareIconURL, spc.ShareURL,
		spc.Sort, spc.IsShow)

	glog.Info("insertSQL:", insertSQL)

	rs, err := db.SQLDB.Exec(insertSQL)
	if nil != err {
		return
	}

	id, err = rs.LastInsertId()
	if nil != err {
		return
	}

	return
}

// ModifySPColumn 修改栏目
func (h *Homepage) ModifySPColumn(spc SPColumn) (id int64, err error) {
	updateSQL := utils.SetSQLFormat(`
		UPDATE 
			gpxj_app.t_homepage_product_stock_recommend 
		SET
			column_id = '{1}',
			stock_name = '{2}',
			stock_code = '{3}',
			select_price = '{4}',
			select_time = '{5}',
			profit_desc = '{6}',
			profit_ratio = '{7}',
			product_name = '{8}',
			product_class = '{9}',
			product_desc = '{10}',
			jump_url = '{11}',
			stock_trend_image_url = '{12}',
			share_title = '{13}',
			share_desc = '{14}',
			share_icon_url = '{15}',
			share_url = '{16}',
			sort = '{17}',
			is_show = '{18}'
		WHERE id = '{0}' ;
		`, spc.ID,
		spc.ColumnID, spc.StockName, spc.StockCode,
		spc.SelectPrice, spc.SelectTime, spc.ProfitDesc, spc.ProfitRatio,
		spc.ProductName, spc.ProductClass, spc.ProductDesc,
		spc.JumpURL, spc.StockTrendImageURL,
		spc.ShareTitle, spc.ShareDesc, spc.ShareIconURL, spc.ShareURL,
		spc.Sort, spc.IsShow)

	glog.Info("updateSQL:", updateSQL)

	stmt, err := db.SQLDB.Prepare(updateSQL)
	if nil != err {
		return
	}

	rs, err := stmt.Exec()
	if nil != err {
		return
	}

	id, err = rs.RowsAffected()
	if nil != err {
		return
	}

	return
}

// DropSPColumn 删除荐股产品栏目
func (h *Homepage) DropSPColumn(spc SPColumn) (id int64, err error) {

	deleteSQL := utils.SetSQLFormat(`DELETE FROM gpxj_app.t_homepage_product_stock_recommend WHERE id ='{0}'`, spc.ID)
	glog.Info("deleteSQL:", deleteSQL)

	rs, err := db.SQLDB.Exec(deleteSQL)
	if nil != err {
		return
	}
	id, err = rs.RowsAffected()
	if nil != err {
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
	Content        string `json:"content"`
	JumpURL        string `json:"jump_url"`
	IsShow         int    `json:"is_show"`

	ShareTitle   string `json:"share_title"`
	ShareIconURL string `json:"share_icon_url"`
	ShareURL     string `json:"share_url"`
	ShareDesc    string `json:"share_desc"`
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

// AddArticleColumn 新增文章栏目
func (h *Homepage) AddArticleColumn(ac ArticleColumn) (id int64, err error) {

	insertSQL := utils.SetSQLFormat(`
	INSERT INTO gpxj_app.t_homepage_article (
		column_id,
		icon_url,nickname,content,jump_url,
		share_title,share_desc,share_icon_url,share_url,
		is_show) 
	  VALUES
		(
		  '{0}',
		  '{1}','{2}','{3}','{4}',
		  '{5}','{6}','{7}','{8}',
		  '{9}') ;
	  
	 `, ac.ColumnID,
		ac.IconURL, ac.Nickname, ac.Content, ac.JumpURL,
		ac.ShareTitle, ac.ShareDesc, ac.ShareIconURL, ac.ShareURL,
		ac.IsShow)

	glog.Info("insertSQL:", insertSQL)

	rs, err := db.SQLDB.Exec(insertSQL)
	if nil != err {
		return
	}

	id, err = rs.LastInsertId()
	if nil != err {
		return
	}

	return
}

// ModifyArticleColumn 修改文章栏目
func (h *Homepage) ModifyArticleColumn(ac ArticleColumn) (id int64, err error) {
	updateSQL := utils.SetSQLFormat(`
	UPDATE 
		gpxj_app.t_homepage_article 
	SET
		column_id = '{1}',
		icon_url = '{2}',
		nickname = '{3}',
		content = '{4}',
		jump_url = '{5}',
		share_title = '{6}',
		share_desc = '{7}',
		share_icon_url = '{8}',
		share_url = '{9}',
		is_show = '{10}'
	WHERE id = '{0}';
	`, ac.ID,
		ac.ColumnID,
		ac.IconURL, ac.Nickname, ac.Content, ac.JumpURL,
		ac.ShareTitle, ac.ShareDesc, ac.ShareIconURL, ac.ShareURL,
		ac.IsShow)

	glog.Info("updateSQL:", updateSQL)

	stmt, err := db.SQLDB.Prepare(updateSQL)
	if nil != err {
		return
	}

	rs, err := stmt.Exec()
	if nil != err {
		return
	}

	id, err = rs.RowsAffected()
	if nil != err {
		return
	}

	return
}

// DropArticleColumn 删除文章栏目
func (h *Homepage) DropArticleColumn(ac ArticleColumn) (id int64, err error) {

	deleteSQL := utils.SetSQLFormat(`DELETE FROM gpxj_app.t_homepage_article WHERE id ='{0}'`, ac.ID)
	glog.Info("deleteSQL:", deleteSQL)

	rs, err := db.SQLDB.Exec(deleteSQL)
	if nil != err {
		return
	}
	id, err = rs.RowsAffected()
	if nil != err {
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
