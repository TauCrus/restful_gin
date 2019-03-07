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

	rows, err := db.SqlDB.Query(querySQL)
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

	rows, err := db.SqlDB.Query(querySQL)
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
