package models

import (
	db "restful_gin/database"
	"restful_gin/utils"
)

// Service 服务
type Service struct {
	ID           int    `json:"id"`
	UserID       string `json:"userid"`
	ProductID    int    `json:"prodict_id"`
	ProductName  string `json:"product_name"`
	ProductClass string `json:"product_class"`
	Enable       int    `json:"enable"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	CreateTime   string `json:"create_time"`
}

// AddService 添加服务
func (s *Service) AddService() (id int64, err error) {
	querySQL := utils.SetSQLFormat(`
		SELECT COUNT(1) 
		FROM gpxj_app.t_service
		WHERE userid = '{0}' AND product_class ='{1}' AND enable = 1
		`, s.UserID, s.ProductClass)

	var cnt int
	err = db.SQLDB.QueryRow(querySQL).Scan(&cnt)
	if nil != err {
		return
	}
	if cnt > 0 {
		return
	}

	rs, err := db.SQLDB.Exec(`
		INSERT INTO 
		gpxj_app.t_service(userid,product_id,product_class,enable,start_date,end_date,create_time) 
		VALUES
		(?,0,?,?,?,?,NOW())
		`, s.UserID, s.ProductClass, s.Enable, s.StartDate, s.EndDate)
	if nil != err {
		return
	}

	id, err = rs.LastInsertId()
	if nil != err {
		return
	}

	return
}

// GetServices 查询服务
func (s *Service) GetServices(userid, productClass string) (services []Service, err error) {

	services = make([]Service, 0)

	querySQL := utils.SetSQLFormat(`
	SELECT a.id,a.userid,a.enable,
		a.product_id,a.product_class,b.product_name,
		a.create_time,a.start_date,a.end_date
	FROM gpxj_app.t_service a
	LEFT JOIN gpxj_app.t_product b ON a.product_class = b.product_class
	WHERE 1
	`)
	if "" != userid {
		querySQL = utils.SetSQLFormat(`{0} AND a.userid like '%{1}%'`, querySQL, userid)
	}

	if "" != productClass {
		querySQL = utils.SetSQLFormat(`{0} AND a.product_class = '{1}'`, querySQL, productClass)
	}

	querySQL = utils.SetSQLFormat(`{0} ORDER BY a.create_time DESC`, querySQL)

	// log.Println("querySQL:", querySQL)

	rows, err := db.SQLDB.Query(querySQL)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var service Service
		rows.Scan(&service.ID, &service.UserID, &service.Enable,
			&service.ProductID, &service.ProductClass, &service.ProductName,
			&service.CreateTime, &service.StartDate, &service.EndDate)

		services = append(services, service)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}

// ModifyService 修改服务
func (s *Service) ModifyService() (id int64, err error) {

	stmt, err := db.SQLDB.Prepare(`		
		UPDATE gpxj_app.t_service
		SET product_class= ?,
			start_date = ?,
			end_date= ?,
			enable = ?
		WHERE id = ?`)
	if nil != err {
		return
	}

	rs, err := stmt.Exec(s.ProductClass, s.StartDate, s.EndDate, s.Enable, s.ID)
	if nil != err {
		return
	}
	id, err = rs.RowsAffected()
	if nil != err {
		return
	}

	return
}

// DropService 删除服务
func (s *Service) DropService() (id int64, err error) {

	rs, err := db.SQLDB.Exec(`DELETE FROM gpxj_app.t_service WHERE id = ?`, s.ID)
	if nil != err {
		return
	}
	id, err = rs.RowsAffected()
	if nil != err {
		return
	}
	return
}
