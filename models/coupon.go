package models

import (
	"log"
	db "restful_gin/database"
	"restful_gin/utils"
)

// Coupon 优惠券
type Coupon struct {
}

// CouponData 优惠券数据
type CouponData struct {
	ID         string  `json:"id"`
	CouponID   int     `json:"coupon_id"`
	CouponName string  `json:"coupon_name"`
	CouponDesc string  `json:"coupon_desc"`
	Type       int     `json:"type"`
	Enable     int     `json:"enable"`
	DaysAdd    int     `json:"days_add"`
	AmountSub  float64 `json:"amount_sub"`
	MinPayment float64 `json:"min_payment"`
	StartTime  string  `json:"start_time"`
	EndTime    string  `json:"end_time"`
	ExpType    int     `json:"exp_type"`
	ExpAddDays int     `json:"exp_add_days"`
	ExpEndDate string  `json:"exp_end_date"`
	Discount   float64 `json:"discount"`
	Activity   string  `json:"activity"`
}

// GetCoupons 查询优惠券
func (c *Coupon) GetCoupons() (cdList []CouponData, err error) {

	cdList = make([]CouponData, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT 
			c.id,
			c.coupon_id,
			c.coupon_name,
			c.coupon_desc,
			c.type,
			c.enable,
			c.days_add,
			c.amount_sub*0.01,
			c.min_payment*0.01,
			c.start_time,
			c.end_time,
			c.exp_type,
			c.exp_add_days ,
			IFNULL(c.exp_end_date,""),
			c.discount * 10,
			IFNULL(caa.activity,"")
		FROM t_coupons c
		LEFT JOIN (SELECT coupon_id AS cid, activity FROM gpxj_app.t_coupon_activity_available) caa ON c.coupon_id = caa.cid
		WHERE 1
	`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY c.id  DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var cd CouponData
		rows.Scan(&cd.ID, &cd.CouponID, &cd.CouponName, &cd.CouponDesc,
			&cd.Type, &cd.Enable, &cd.DaysAdd, &cd.AmountSub, &cd.MinPayment,
			&cd.StartTime, &cd.EndTime, &cd.ExpType, &cd.ExpAddDays, &cd.ExpEndDate,
			&cd.Discount, &cd.Activity)

		cdList = append(cdList, cd)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}

// CouponRecord 优惠券领取记录
type CouponRecord struct {
	ID          int    `json:"id"`
	CouponID    int    `json:"coupon_id"`
	Userid      string `json:"userid"`
	Source      string `json:"source"`
	Channel     int    `json:"channel"`
	Activity    string `json:"activity"`
	Consume     int    `json:"consume"`
	ExpiredDate string `json:"expired_date"`
	OrderID     string `json:"order_id"`
}

// GetCouponRecord 查询优惠券领取记录
func (c *Coupon) GetCouponRecord() (coList []CouponRecord, err error) {

	coList = make([]CouponRecord, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id,coupon_id,userid,
			IFNULL(source,""),
			channel,activity,consume,
			expired_date,
			IFNULL(order_id,"") 
		FROM gpxj_app.t_coupon_obtain 
		WHERE 1
	`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id  DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var co CouponRecord
		rows.Scan(&co.ID, &co.CouponID, &co.Userid, &co.Source,
			&co.Channel, &co.Activity, &co.Consume,
			&co.ExpiredDate, &co.OrderID)

		coList = append(coList, co)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}

// CouponActivity 优惠券活动
type CouponActivity struct {
	ID           int    `json:"id"`
	Activity     string `json:"activity"`
	ActivityDesc string `json:"activity_desc"`
	Uptime       string `json:"uptime"`
	EndTime      string `json:"end_time"`
	IsShow       int    `json:"is_show"`
}

// GetCouponActivity 查询优惠券活动
func (c *Coupon) GetCouponActivity() (caList []CouponActivity, err error) {
	caList = make([]CouponActivity, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id,activity,activity_desc,
			uptime,end_time,is_show 
		FROM t_coupon_activity_setting
		WHERE 1
	`)

	querySQL = utils.SetSQLFormat(`{0} ORDER BY id  DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var ca CouponActivity
		rows.Scan(&ca.ID, &ca.Activity, &ca.ActivityDesc,
			&ca.Uptime, &ca.EndTime, &ca.IsShow)

		caList = append(caList, ca)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}
