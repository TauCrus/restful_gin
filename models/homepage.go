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
func (h *Homepage) GetColumns(columnTypeID, columnTypeIDs, columnName, status string) (columns []Column, err error) {
	columns = make([]Column, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id,
			column_type_id,column_type_name,
			column_name ,jump_url,jump_id,
			sort,is_show,in_review
		FROM  gpxj_app.t_column 
		WHERE 1
	`)

	if "0" != columnTypeID {
		querySQL = utils.SetSQLFormat(`{0} AND column_type_id = '{1}'`, querySQL, columnTypeID)
	}
	if "" != columnTypeIDs {
		querySQL = utils.SetSQLFormat(`{0} AND column_type_id IN ({1})`, querySQL, columnTypeIDs)

	}
	if "" != columnName {
		querySQL = utils.SetSQLFormat(`{0} AND column_name like '%{1}%'`, querySQL, columnName)
	}

	if "" != status {
		if "1" == status {
			querySQL = utils.SetSQLFormat(`{0} AND is_show = 1`, querySQL)
		} else if "2" == status {
			querySQL = utils.SetSQLFormat(`{0} AND is_show = 0`, querySQL)
		} else if "3" == status {
			querySQL = utils.SetSQLFormat(`{0} AND in_review = 1`, querySQL)
		}
	}
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
		WHERE id = '{0}' ;
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

// ColumnType 栏目类型
type ColumnType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetColumnTypes 获取栏目分类
func (h *Homepage) GetColumnTypes() (cts []ColumnType, err error) {
	cts = make([]ColumnType, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id,name
		FROM gpxj_app.t_column_type
		WHERE id <> 1
		ORDER BY id
	`)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}
	for rows.Next() {
		var ct ColumnType
		rows.Scan(&ct.ID, &ct.Name)

		cts = append(cts, ct)
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
	ColumnTypeID   int     `json:"column_type_id"`
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
func (h *Homepage) GetProductColumns(keyword, columnID, status string) (pcs []ProductColumn, err error) {

	pcs = make([]ProductColumn, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT  hpc.id,
			hpc.column_id,
			c.column_name,
			c.column_type_id,
			c.column_type_name,
			hpc.image_url ,
			hpc.jump_url ,
			hpc.product_name,
			hpc.product_class,
			hpc.product_desc ,
			hpc.subscribe_num ,
			hpc.price,
			hpc.sort,
			hpc.is_show,
			IFNULL(hpc.share_title,''),
			IFNULL(hpc.share_icon_url,''),
			IFNULL(hpc.share_url,''),
			IFNULL(hpc.share_desc,'')
		FROM t_homepage_product_recommend hpc
		LEFT JOIN  t_column c ON c.id = hpc.column_id
		WHERE 1
	`)

	if "" != keyword {
		querySQL = utils.SetSQLFormat(`{0} AND CONCAT(c.column_name,c.column_type_name,hpc.product_name,hpc.product_class) like '%{1}%'`, querySQL, keyword)
	}

	if "0" != columnID {
		querySQL = utils.SetSQLFormat(`{0} AND hpc.column_id = '{1}'`, querySQL, columnID)
	}

	if "" != status {
		if "1" == status {
			querySQL = utils.SetSQLFormat(`{0} AND hpc.is_show = 1`, querySQL)
		} else if "2" == status {
			querySQL = utils.SetSQLFormat(`{0} AND hpc.is_show = 0`, querySQL)
		}
	}

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}
	for rows.Next() {
		var pc ProductColumn
		rows.Scan(&pc.ID, &pc.ColumnID, &pc.ColumnName, &pc.ColumnTypeID, &pc.ColumnTypeName,
			&pc.ImageURL, &pc.JumpURL, &pc.ProductName, &pc.ProductClass, &pc.ProductDesc,
			&pc.SubscribeNum, &pc.Price, &pc.Sort, &pc.IsShow,
			&pc.ShareTitle, &pc.ShareIconURL, &pc.ShareURL, &pc.ShareDesc)

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
	ColumnTypeID       int     `json:"column_type_id"`
	ColumnTypeName     string  `json:"column_type_name"`
	StockName          string  `json:"stock_name"`
	StockCode          string  `json:"stock_code"`
	SelectPrice        int     `json:"select_price"`
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
func (h *Homepage) GetSPColumns(keyword, columnID, status string) (spcs []SPColumn, err error) {

	spcs = make([]SPColumn, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT 
			hpsr.id,
			hpsr.column_id,
			c.column_name,
			c.column_type_id,
			c.column_type_name,
			hpsr.stock_name,
			hpsr.stock_code,
			hpsr.product_name,
			hpsr.product_class,
			hpsr.product_desc,
			hpsr.select_price,
			hpsr.select_time,
			hpsr.profit_desc,
			hpsr.profit_ratio,
			hpsr.jump_url,
			hpsr.stock_trend_image_url,
			hpsr.sort,
			hpsr.is_show,
			IFNULL(hpsr.share_title,''),
			IFNULL(hpsr.share_icon_url,''),
			IFNULL(hpsr.share_url,''),
			IFNULL(hpsr.share_desc,'')
		FROM gpxj_app.t_homepage_product_stock_recommend hpsr
		LEFT JOIN gpxj_app.t_column c ON hpsr.column_id = c.id
		WHERE 1
	`)

	if "" != keyword {
		querySQL = utils.SetSQLFormat(`{0} AND CONCAT(c.column_name,c.column_type_name,hpsr.stock_name,hpsr.stock_code,hpsr.product_name,hpsr.product_class) like '%{1}%'`, querySQL, keyword)
	}

	if "0" != columnID {
		querySQL = utils.SetSQLFormat(`{0} AND hpsr.column_id = '{1}'`, querySQL, columnID)
	}

	if "" != status {
		if "1" == status {
			querySQL = utils.SetSQLFormat(`{0} AND hpsr.is_show = 1`, querySQL)
		} else if "2" == status {
			querySQL = utils.SetSQLFormat(`{0} AND hpsr.is_show = 0`, querySQL)
		}
	}

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}
	for rows.Next() {
		var spc SPColumn
		rows.Scan(&spc.ID, &spc.ColumnID, &spc.ColumnName, &spc.ColumnTypeID, &spc.ColumnTypeName,
			&spc.StockName, &spc.StockCode, &spc.ProductName, &spc.ProductClass, &spc.ProductDesc,
			&spc.SelectPrice, &spc.SelectTime, &spc.ProfitDesc, &spc.ProfitRatio,
			&spc.JumpURL, &spc.StockTrendImageURL, &spc.Sort, &spc.IsShow,
			&spc.ShareTitle, &spc.ShareIconURL, &spc.ShareURL, &spc.ShareDesc)

		spcs = append(spcs, spc)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}

// AddSPColumn 新增荐股产品栏目
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

// ModifySPColumn 修改荐股产品栏目
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
	ColumnTypeID   int    `json:"column_type_id"`
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
func (h *Homepage) GetArticleColumns(columnID, status string) (acs []ArticleColumn, err error) {

	acs = make([]ArticleColumn, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT 
			ha.id,
			ha.column_id ,
			c.column_name,
			c.column_type_id,
			c.column_type_name,
			ha.icon_url ,
			ha.nickname,
			ha.content,
			ha.jump_url,
			ha.is_show,
			IFNULL(ha.share_title,''),
			IFNULL(ha.share_icon_url,''),
			IFNULL(ha.share_url,''),
			IFNULL(ha.share_desc,'')
		FROM t_homepage_article ha
		LEFT JOIN t_column c ON ha.column_id = c.id
		WHERE 1
	`)

	if "0" != columnID {
		querySQL = utils.SetSQLFormat(`{0} AND ha.column_id = '{1}'`, querySQL, columnID)
	}

	if "" != status {
		if "1" == status {
			querySQL = utils.SetSQLFormat(`{0} AND ha.is_show = 1`, querySQL)
		} else if "2" == status {
			querySQL = utils.SetSQLFormat(`{0} AND ha.is_show = 0`, querySQL)
		}
	}

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}
	for rows.Next() {
		var ac ArticleColumn
		rows.Scan(&ac.ID, &ac.ColumnID, &ac.ColumnName, &ac.ColumnTypeID, &ac.ColumnTypeName,
			&ac.IconURL, &ac.Nickname, &ac.Content, &ac.JumpURL, &ac.IsShow,
			&ac.ShareTitle, &ac.ShareIconURL, &ac.ShareURL, &ac.ShareDesc)

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

	ShareTitle   string `json:"share_title"`
	ShareIconURL string `json:"share_icon_url"`
	ShareURL     string `json:"share_url"`
	ShareDesc    string `json:"share_desc"`
}

// GetShortCutMenus 查询首页快捷菜单
func (h *Homepage) GetShortCutMenus(keyword, status string) (scms []ShortCutMenu, err error) {

	scms = make([]ShortCutMenu, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id, title,icon_url, jump_url,
			is_new, in_review, product_class,
			IFNULL(share_title,''),
			IFNULL(share_icon_url,''),
			IFNULL(share_url,''),
			IFNULL(share_desc,'')
		FROM gpxj_app.t_homepage_product_menu
		WHERE 1
		`)

	if "" != keyword {
		querySQL = utils.SetSQLFormat(`{0} AND title like '%{1}%'`, querySQL, keyword)
	}

	if "" != status {
		if "1" == status {
			querySQL = utils.SetSQLFormat(`{0} AND in_review = 1`, querySQL)
		} else if "2" == status {
			querySQL = utils.SetSQLFormat(`{0} AND is_new = 1`, querySQL)
		}
	}

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var scm ShortCutMenu
		rows.Scan(&scm.ID, &scm.Title, &scm.IconURL, &scm.JumpURL,
			&scm.IsNew, &scm.InReview, &scm.ProductClass,
			&scm.ShareTitle, &scm.ShareIconURL, &scm.ShareURL, &scm.ShareDesc)

		scms = append(scms, scm)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}

// AddShortCutMenu 新增首页快捷菜单
func (h *Homepage) AddShortCutMenu(scm ShortCutMenu) (id int64, err error) {

	insertSQL := utils.SetSQLFormat(`
	INSERT INTO gpxj_app.t_homepage_product_menu (
		title,icon_url,jump_url,
		share_title,share_desc,share_icon_url,share_url,
		is_new,in_review,product_class) 
	  VALUES
		(
		  '{0}','{1}','{2}',
		  '{3}','{4}','{5}','{6}',
		  '{7}','{8}','{9}');
		`, scm.Title, scm.IconURL, scm.JumpURL,
		scm.ShareTitle, scm.ShareDesc, scm.ShareIconURL, scm.ShareURL,
		scm.IsNew, scm.InReview, scm.ProductClass)

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

// ModifyShortCutMenu 修改首页快捷菜单
func (h *Homepage) ModifyShortCutMenu(scm ShortCutMenu) (id int64, err error) {
	updateSQL := utils.SetSQLFormat(`
		UPDATE 
			gpxj_app.t_homepage_product_menu 
		SET
			title = '{1}',
			icon_url = '{2}',
			jump_url = '{3}',
			share_title = '{4}',
			share_desc = '{5}',
			share_icon_url = '{6}',
			share_url = '{7}',
			is_new = '{8}',
			in_review = '{9}',
			product_class = '{10}' 
		WHERE id = '{0}';
		`, scm.ID,
		scm.Title, scm.IconURL, scm.JumpURL,
		scm.ShareTitle, scm.ShareDesc, scm.ShareIconURL, scm.ShareURL,
		scm.IsNew, scm.InReview, scm.ProductClass)

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

// DropShortCutMenu 删除首页快捷菜单
func (h *Homepage) DropShortCutMenu(scm ShortCutMenu) (id int64, err error) {

	deleteSQL := utils.SetSQLFormat(`DELETE FROM gpxj_app.t_homepage_product_menu WHERE id ='{0}'`, scm.ID)
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

// ProductClassify 产品分类
type ProductClassify struct {
	ID               int    `json:"id"`
	ProductTypeName  string `json:"product_type_name"`
	ProductTypeDesc  string `json:"product_type_desc"`
	ProductClass     string `json:"product_class"`
	SmallIconURL     string `json:"small_icon_url"`
	ListIconURL      string `json:"list_icon_url"`
	InnerJumpURL     string `json:"inner_jump_url"`
	HasNew           int    `json:"has_new"`
	InReview         int    `json:"in_review"`
	IsShow           int    `json:"is_show"`
	Sort             int    `json:"sort"`
	ApplePayPID      string `json:"apple_pay_pid"`
	ApplePayPrice    string `json:"apple_pay_price"`
	ApplePayVIPPrice string `json:"apple_pay_vip_price"`
}

// GetProductClassifys 查询首页产品分类
func (h *Homepage) GetProductClassifys(keyword, status string) (pcs []ProductClassify, err error) {
	pcs = make([]ProductClassify, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id,
			product_type_name, product_type_desc, product_class,
			IFNULL(small_icon_url,""),
			IFNULL(list_icon_url,""), 
			IFNULL(inner_jump_url,""),
			has_new,in_review, is_show, sort,
			IFNULL(apple_pay_pid,''),
			IFNULL(apple_pay_price,''),
			IFNULL(apple_pay_vip_price,'')
		FROM gpxj_app.t_product_type 
		WHERE 1
	`)

	if "" != keyword {
		querySQL = utils.SetSQLFormat(`{0} AND CONCAT(product_type_name,product_type_desc,product_class) like '%{1}%'`, querySQL, keyword)
	}

	if "" != status {
		if "1" == status {
			querySQL = utils.SetSQLFormat(`{0} AND is_show = 1`, querySQL)
		} else if "2" == status {
			querySQL = utils.SetSQLFormat(`{0} AND is_show = 0`, querySQL)
		} else if "3" == status {
			querySQL = utils.SetSQLFormat(`{0} AND in_review = 1`, querySQL)
		}
	}

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
			&pc.HasNew, &pc.InReview, &pc.IsShow, &pc.Sort,
			&pc.ApplePayPID, &pc.ApplePayPrice, &pc.ApplePayVIPPrice)

		pcs = append(pcs, pc)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}

// AddProductClassify 新增首页产品分类
func (h *Homepage) AddProductClassify(pc ProductClassify) (id int64, err error) {

	insertSQL := utils.SetSQLFormat(`
		INSERT INTO gpxj_app.t_product_type (
			product_type_name,product_type_desc,
			small_icon_url,list_icon_url,inner_jump_url,
			product_class,has_new,in_review,is_show,sort,
			apple_pay_pid,apple_pay_price,apple_pay_vip_price
		) 
	 	 VALUES
		(
		  '{0}','{1}',
		  '{2}','{3}','{4}',
		  '{5}','{6}','{7}','{8}','{9}',
		  '{10}','{11}','{12}') ;
		`, pc.ProductTypeName, pc.ProductTypeDesc,
		pc.SmallIconURL, pc.ListIconURL, pc.InnerJumpURL,
		pc.ProductClass, pc.HasNew, pc.InReview, pc.IsShow, pc.Sort,
		pc.ApplePayPID, pc.ApplePayPrice, pc.ApplePayVIPPrice)

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

// ModifyProductClassify 修改首页产品分类
func (h *Homepage) ModifyProductClassify(pc ProductClassify) (id int64, err error) {
	updateSQL := utils.SetSQLFormat(`
		UPDATE 
			gpxj_app.t_product_type 
		SET
			product_type_name = '{1}',
			product_type_desc = '{2}',
			small_icon_url = '{3}',
			list_icon_url = '{4}',
			inner_jump_url = '{5}',
			product_class = '{6}',
			has_new = '{7}',
			in_review = '{8}',
			is_show = '{9}',
			sort = '{10}',
			apple_pay_pid = '{11}',
			apple_pay_price = '{12}',
			apple_pay_vip_price = '{13}'
		WHERE id = '{0}' ;
		`, pc.ID,
		pc.ProductTypeName, pc.ProductTypeDesc,
		pc.SmallIconURL, pc.ListIconURL, pc.InnerJumpURL,
		pc.ProductClass, pc.HasNew, pc.InReview, pc.IsShow, pc.Sort,
		pc.ApplePayPID, pc.ApplePayPrice, pc.ApplePayVIPPrice)

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

// DropProductClassify 删除首页产品分类
func (h *Homepage) DropProductClassify(pc ProductClassify) (id int64, err error) {

	deleteSQL := utils.SetSQLFormat(`DELETE FROM gpxj_app.t_product_type WHERE id ='{0}'`, pc.ID)
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

// ActivityMarketing 活动营销
type ActivityMarketing struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	ImageURL      string `json:"image_url"`
	IconURL       string `json:"icon_url"`
	JumpURL       string `json:"jump_url"`
	Uptime        string `json:"uptime"`
	ActivityType  int    `json:"activity_type"`
	IsShow        int    `json:"is_show"`
	MarketingType string `json:"marketing_type"`
	ShareTitle    string `json:"share_title"`
	ShareIconURL  string `json:"share_icon_url"`
	ShareURL      string `json:"share_url"`
	ShareDesc     string `json:"share_desc"`
}

// GetActivityMarketings 查询活动营销
func (h *Homepage) GetActivityMarketings(mType int, keyword, status string) (ams []ActivityMarketing, err error) {
	ams = make([]ActivityMarketing, 0)
	querySQL := utils.SetSQLFormat(`
		SELECT id,title,
			IFNULL(image_url,""),
			IFNULL(icon_url,""), 
			IFNULL(jump_url,""),
			uptime,activity_type,is_show,
			IFNULL(share_title,''),
			IFNULL(share_icon_url,''),
			IFNULL(share_url,''),
			IFNULL(share_desc,'')
		FROM gpxj_app.t_activity_marketing
		WHERE 1
	`)

	if 1 == mType {
		querySQL = utils.SetSQLFormat(`{0} AND marketing_type = 'suspension'`, querySQL)
	} else if 2 == mType {
		querySQL = utils.SetSQLFormat(`{0} AND marketing_type = 'popups'`, querySQL)
	}

	if "" != keyword {
		querySQL = utils.SetSQLFormat(`{0} AND title like '%{1}%'`, querySQL, keyword)
	}

	if "" != status {
		if "1" == status {
			querySQL = utils.SetSQLFormat(`{0} AND is_show = 1`, querySQL)
		} else if "2" == status {
			querySQL = utils.SetSQLFormat(`{0} AND is_show = 0`, querySQL)
		}
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
			&am.Uptime, &am.ActivityType, &am.IsShow,
			&am.ShareTitle, &am.ShareIconURL, &am.ShareURL, &am.ShareDesc)

		ams = append(ams, am)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}

// AddActivityMarketing 新增营销活动
func (h *Homepage) AddActivityMarketing(am ActivityMarketing) (id int64, err error) {

	insertSQL := utils.SetSQLFormat(`
	INSERT INTO gpxj_app.t_activity_marketing (
		title,image_url,icon_url,jump_url,
		share_title,share_desc,share_icon_url,share_url,
		uptime,activity_type,is_show,marketing_type) 
	  VALUES
		(
		  '{0}','{1}','{2}','{3}',
		  '{4}','{5}','{6}','{7}',
		  '{8}','{9}','{10}','{11}') ;
		`, am.Title, am.ImageURL, am.IconURL, am.JumpURL,
		am.ShareTitle, am.ShareDesc, am.ShareIconURL, am.ShareURL,
		am.Uptime, am.ActivityType, am.IsShow, am.MarketingType)

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

// ModifyActivityMarketing 修改营销活动
func (h *Homepage) ModifyActivityMarketing(am ActivityMarketing) (id int64, err error) {
	updateSQL := utils.SetSQLFormat(`
		UPDATE 
			gpxj_app.t_activity_marketing 
		SET
			title = '{1}',
			image_url = '{2}',
			icon_url = '{3}',
			jump_url = '{4}',
			share_title = '{5}',
			share_desc = '{6}',
			share_icon_url = '{7}',
			share_url = '{8}',
			uptime = '{9}',
			activity_type = '{10}',
			is_show = '{11}',
			marketing_type = '{12}' 
		WHERE id = '{0}' ;
		`, am.ID,
		am.Title, am.ImageURL, am.IconURL, am.JumpURL,
		am.ShareTitle, am.ShareDesc, am.ShareIconURL, am.ShareURL,
		am.Uptime, am.ActivityType, am.IsShow, am.MarketingType)

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

// DropActivityMarketing 删除营销活动
func (h *Homepage) DropActivityMarketing(am ActivityMarketing) (id int64, err error) {

	deleteSQL := utils.SetSQLFormat(`DELETE FROM gpxj_app.t_activity_marketing WHERE id ='{0}'`, am.ID)
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
