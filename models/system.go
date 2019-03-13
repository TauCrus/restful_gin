package models

import db "restful_gin/database"

// System 系统
type System struct{}

// Review 审核数据
type Review struct {
	ID       int    `json:"id"`
	Channel  int    `json:"channel"`
	Version  string `json:"version"`
	InReview int    `json:"inReview"`
	Memo     string `json:"memo"`
}

// SysReviewsQuery 查询所有用户
func (s *System) SysReviewsQuery() (data []Review, err error) {

	data = make([]Review, 0)

	rows, err := db.SQLDB.Query(`
		SELECT id,channel,version,in_review,memo
		FROM gpxj_app.t_release_review
	`)
	if nil != err {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r Review
		rows.Scan(&r.ID, &r.Channel, &r.Version, &r.InReview, &r.Memo)

		data = append(data, r)
	}
	if err = rows.Err(); nil != err {
		return
	}
	return
}
