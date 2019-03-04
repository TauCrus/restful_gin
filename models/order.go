package models

import db "restful_gin/database"

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
func (o *Order) QueryOrders() (orders []Order, err error) {

	orders = make([]Order, 0)

	rows, err := db.SqlDB.Query(`
	SELECT order_id, pay_way, userid,
			product_id,product_class,
			quantity,total_fee/100, 
			DATE_FORMAT(create_time,'%Y年%m月%d日'),
			DATE_FORMAT(start_date,'%Y年%m月%d日'),
			DATE_FORMAT(end_date,'%Y年%m月%d日')
	FROM gpxj_app.t_order_total 	
	`)

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
