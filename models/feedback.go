package models

import (
	"log"
	db "restful_gin/database"
	"restful_gin/utils"
)

// Feedback 反馈
type Feedback struct {
	ID          int    `json:"id"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Content     string `json:"content"`
	OperateTime string `json:"operate_time"`
}

// GetFeedbacks 查询服务
func (f *Feedback) GetFeedbacks(userid, email string) (feedbacks []Feedback, err error) {

	feedbacks = make([]Feedback, 0)

	querySQL := utils.SetSQLFormat(`
		SELECT id,phone,email,content,time 
		FROM gpxj_app.t_feedback
		WHERE 1 
	`)
	if "" != userid {
		querySQL = utils.SetSQLFormat(`{0} AND phone = '{1}'`, querySQL, userid)
	}

	if "" != email {
		querySQL = utils.SetSQLFormat(`{0} AND email = '{1}'`, querySQL, email)
	}

	querySQL = utils.SetSQLFormat(`{0} ORDER BY time DESC`, querySQL)

	log.Println("querySQL:", querySQL)

	rows, err := db.SqlDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var feedback Feedback
		rows.Scan(&feedback.ID, &feedback.Phone, &feedback.Email, &feedback.Content, &feedback.OperateTime)

		feedbacks = append(feedbacks, feedback)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}

// DropFeedback 删除服务
func (f *Feedback) DropFeedback() (id int64, err error) {

	rs, err := db.SqlDB.Exec(`DELETE FROM gpxj_app.t_feedback WHERE id = ?`, f.ID)
	if nil != err {
		return
	}
	id, err = rs.RowsAffected()
	if nil != err {
		return
	}
	return
}
