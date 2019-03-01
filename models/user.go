package models

import db "restful_gin/database"

// User 用户
type User struct {
	ID       string `json:"id"`
	RealName string `json:"realname"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// GetSysUsers 查询所有用户
func (u *User) GetSysUsers() (users []User, err error) {

	users = make([]User, 0)

	rows, err := db.SqlDB.Query(`
		SELECT id ,realname,username,email,phone
		FROM gpxj_app.t_admin_user
	`)
	if nil != err {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.RealName, &user.UserName, &user.Email, &user.Phone)

		users = append(users, user)
	}
	if err = rows.Err(); nil != err {
		return
	}
	return
}
