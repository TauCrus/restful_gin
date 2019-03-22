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

	querySQL = utils.SetSQLFormat(`{0} ORDER BY regist_date DESC`, querySQL)
	// 排序

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
