package models

import (
	"database/sql"
	db "restful_gin/database"
	"restful_gin/utils"

	"github.com/golang/glog"
)

// AdminUser 系统用户
type AdminUser struct {
	ID       string `json:"id"`
	RealName string `json:"realname"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// GetAdminUsers 查询所有用户
func (u *User) GetAdminUsers() (users []AdminUser, err error) {

	users = make([]AdminUser, 0)

	rows, err := db.SQLDB.Query(`
		SELECT id ,realname,username,email,phone
		FROM gpxj_app.t_admin_user
	`)
	if nil != err {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user AdminUser
		rows.Scan(&user.ID, &user.RealName, &user.UserName, &user.Email, &user.Phone)

		users = append(users, user)
	}
	if err = rows.Err(); nil != err {
		return
	}
	return
}

// User 用户
type User struct {
	UserID     string `json:"userid"`
	ImageURL   string `json:"image_url"`
	Birthday   string `json:"birthday"`
	IsVIP      int    `json:"is_vip"`
	VIPEndDate string `json:"vip_end_date"`
	Version    string `json:"version"`
	NickName   string `json:"nickname"`
	RegistTime string `json:"regist_time"`
}

// GetUsers 查询所有用户
func (u *User) GetUsers() (users []User, err error) {

	users = make([]User, 0)

	rows, err := db.SQLDB.Query(`		
		SELECT a.userid,
			IFNULL(a.image_url,''),
			IFNULL(a.birthday,''),
			a.is_vip,
			IFNULL(a.vip_end_date,''),
			a.version,
			b.nickname,b.registtime	
		FROM t_user_new a
		LEFT JOIN t_user_base b ON a.userid = b.userid
	`)
	if nil != err {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		rows.Scan(&user.UserID, &user.ImageURL, &user.Birthday, &user.IsVIP, &user.VIPEndDate, &user.Version, &user.NickName, &user.RegistTime)

		users = append(users, user)
	}
	if err = rows.Err(); nil != err {
		return
	}
	return
}

// UserLogin 用户登录
func (u *User) UserLogin(username, password string) (cnt int, err error) {

	querySQL := utils.SetSQLFormat(`
	SELECT COUNT(1)
	FROM gpxj_app.t_back_user
	WHERE username = '{0}' AND password = '{1}' 
	`, username, password)

	err = db.SQLDB.QueryRow(querySQL).Scan(&cnt)
	if nil != err {
		return
	}

	return
}

// GetUserInfo 获取用户信息
func (u *User) GetUserInfo(username string) (roles, permissions []string, err error) {

	roles, err = getUserRoles(username)
	if nil != err {
		return
	}

	permissions, err = getUserPermissions(username)
	if nil != err {
		return
	}

	return
}

// 获取用户角色
func getUserRoles(username string) (roles []string, err error) {
	querySQL := utils.SetSQLFormat(`
		SELECT br.code
		FROM gpxj_app.t_back_role br 
		LEFT JOIN gpxj_app.t_back_user_role bur ON br.id = bur.role_id
		LEFT JOIN gpxj_app.t_back_user bu ON bur.user_id = bu.id
		WHERE bu.username = '{0}'
		`, username)

	glog.Info("querySQL:", querySQL)
	rows, err := db.SQLDB.Query(querySQL)
	if nil != err {
		if err != sql.ErrNoRows {
			return
		}
	}
	defer rows.Close()

	for rows.Next() {
		var role string
		rows.Scan(&role)

		roles = append(roles, role)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}

// 获取用户权限
func getUserPermissions(username string) (permissions []string, err error) {
	querySQL := utils.SetSQLFormat(`
		SELECT bm.url	
		FROM gpxj_app.t_back_menu bm
		LEFT JOIN gpxj_app.t_back_role_menu brm ON bm.id = brm.menu_id
		LEFT JOIN gpxj_app.t_back_user_role bur ON brm.role_id = bur.role_id
		LEFT JOIN gpxj_app.t_back_user bu ON bur.user_id = bu.id
		WHERE bu.username = '{0}'
		GROUP BY bm.url
		ORDER BY bm.sort
		`, username)

	glog.Info("querySQL:", querySQL)
	rows, err := db.SQLDB.Query(querySQL)
	if nil != err {
		if err != sql.ErrNoRows {
			return
		}
	}
	defer rows.Close()

	for rows.Next() {
		var tmp string
		rows.Scan(&tmp)

		permissions = append(permissions, tmp)
	}

	if err = rows.Err(); nil != err {
		return
	}

	return
}
