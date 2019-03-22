package models

import (
	db "restful_gin/database"
	"restful_gin/utils"

	"github.com/golang/glog"
)

// Count 统计
type Count struct {
}

// UserRegist 用户注册
type UserRegist struct {
	UserNum    int    `json:"user_num"`
	RegistDate string `json:"regist_date"`
}

// QueryUserRegist 查询用户注册
func (c *Count) QueryUserRegist(startDate, endDate, group string) (urs []UserRegist, err error) {
	urs = make([]UserRegist, 0)

	querySQL := utils.SetSQLFormat(` 
		SELECT COUNT(userid) AS user_num,
			IF('{0}'='day',DATE(registtime),DATE_FORMAT(registtime,'%Y-%m')) AS regist_date
		FROM gpxj_app.t_user_base
		WHERE 1
	`, group)

	if "" != startDate && "" != endDate {
		querySQL = utils.SetSQLFormat(`{0} AND DATE(registtime) BETWEEN '{1}' AND '{2}' `, querySQL, startDate, endDate)
	}
	// 分组
	querySQL = utils.SetSQLFormat(`{0} 	GROUP BY regist_date`, querySQL)

	// 排序
	querySQL = utils.SetSQLFormat(`{0} ORDER BY regist_date DESC`, querySQL)

	glog.Info("querySQL:", querySQL)
	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var ur UserRegist
		rows.Scan(&ur.UserNum, &ur.RegistDate)
		urs = append(urs, ur)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}

//Order 订单
type Order struct {
	OrderID      string  `json:"order_id"`
	PayWay       string  `json:"pay_way"`
	UserID       string  `json:"userid"`
	ProductID    int     `json:"product_id"`
	ProductClass string  `json:"product_class"`
	Quantity     int     `json:"quantity"`
	TotalFee     float64 `json:"total_fee"`
	CreateTime   string  `json:"create_time"`
	StartDate    string  `json:"start_date"`
	EndDate      string  `json:"end_date"`
}

//QueryOrders 查询所有订单
func (c *Count) QueryOrders(startDate, endDate, orderID, payWay, userid, productClass string) (orders []Order, err error) {

	orders = make([]Order, 0)

	querySQL := utils.SetSQLFormat(` 
		SELECT order_id, pay_way, userid,
			product_id,product_class,
			quantity,total_fee/100, 
			DATE_FORMAT(create_time,'%Y年%m月%d日'),
			DATE_FORMAT(start_date,'%Y年%m月%d日'),
			DATE_FORMAT(end_date,'%Y年%m月%d日')
		FROM gpxj_app.t_order_total 
		WHERE 1
	`)
	if "" != startDate && "" != endDate {
		querySQL = utils.SetSQLFormat(`{0} AND DATE(create_time) BETWEEN '{1}' AND '{2}' `, querySQL, startDate, endDate)
	}

	if "" != orderID {
		querySQL = utils.SetSQLFormat(`{0} AND order_id = '{1}' `, querySQL, orderID)
	}
	if "" != payWay {
		querySQL = utils.SetSQLFormat(`{0} AND pay_way = '{1}' `, querySQL, payWay)
	}
	if "" != userid {
		querySQL = utils.SetSQLFormat(`{0} AND userid = '{1}' `, querySQL, userid)
	}
	if "" != productClass {
		querySQL = utils.SetSQLFormat(`{0} AND product_class =  '{1}' `, querySQL, productClass)
	}

	// 排序
	querySQL = utils.SetSQLFormat(`{0} ORDER BY create_time DESC`, querySQL)

	glog.Info("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	if nil != err {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		rows.Scan(&order.OrderID, &order.PayWay, &order.UserID,
			&order.ProductID, &order.ProductClass, &order.Quantity, &order.TotalFee,
			&order.CreateTime, &order.StartDate, &order.EndDate)

		orders = append(orders, order)
	}
	if err = rows.Err(); nil != err {
		return
	}

	return
}
