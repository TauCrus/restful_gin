package models

import db "restful_gin/database"

// 服务管理

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

	rs, err := db.SqlDB.Exec(``)
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
func (s *Service) GetServices() (services []Service, err error) {

	services = make([]Service, 0)

	rows, err := db.SqlDB.Query(`
	SELECT a.id,a.userid,a.enable,
		a.product_id,a.product_class,b.product_name,
		a.create_time,a.start_date,a.end_date
	FROM gpxj_app.t_service a
	LEFT JOIN gpxj_app.t_product b ON a.product_class = b.product_class
	ORDER BY a.create_time DESC
	`)
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var service Service
		rows.Scan(&service.ID, &service.UserID, &service.Enable,
			&service.ProductID, &service.ProductClass, &service.ProductName,
			&service.StartDate, &service.EndDate, &service.CreateTime)

		services = append(services, service)
	}

	if err = rows.Err(); nil != err {
		return
	}
	return
}

// GetService 查询单个服务，更新时候使用
func (s *Service) GetService(id string) (service Service, err error) {

	err = db.SqlDB.QueryRow(`
			SELECT a.id,a.userid,a.enable,
				a.product_id,a.product_class,b.product_name,
				a.create_time,a.start_date,a.end_date
			FROM gpxj_app.t_service a
			LEFT JOIN gpxj_app.t_product b ON a.product_class = b.product_class 
			WHERE a.id = ?`, id).Scan(&service.ID, &service.UserID, &service.Enable,
		&service.ProductID, &service.ProductClass, &service.ProductName,
		&service.StartDate, &service.EndDate, &service.CreateTime)
	if nil != err {
		return
	}

	return
}

// ModifyService 修改服务
func (s *Service) ModifyService() (id int64, err error) {

	return
}

// DropService 删除服务（暂无）
func (s *Service) DropService() (id int64, err error) {

	return
}
