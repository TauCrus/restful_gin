package models

import (
	// "log"
	db "restful_gin/database"
)

// Person 成员
type Person struct {
	ID        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

//AddPerson 新增
func (p *Person) AddPerson() (id int64, err error) {
	rs, err := db.SQLDB.Exec("INSERT INTO person(first_name,last_name) VALUES(?,?)", p.FirstName, p.LastName)
	if nil != err {
		return
	}

	id, err = rs.LastInsertId()
	if nil != err {
		return
	}
	return
}

//GetPersons  查询
func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)

	rows, err := db.SQLDB.Query("SELECT id, first_name, last_name FROM person")
	defer rows.Close()

	if nil != err {
		return
	}

	for rows.Next() {
		var person Person
		rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); nil != err {
		return
	}
	return
}

// GetPerson 查询单条
func (p *Person) GetPerson(id string) (person Person, err error) {

	err = db.SQLDB.QueryRow("SELECT id, first_name, last_name FROM person WHERE id =?", id).Scan(&person.ID, &person.FirstName, &person.LastName)
	if nil != err {
		return
	}

	return
}

//UpdatePerson 更新
func (p *Person) UpdatePerson() (id int64, err error) {

	stmt, err := db.SQLDB.Prepare("UPDATE person SET first_name=?,last_name=? WHERE id=?")
	if nil != err {
		return
	}
	rs, err := stmt.Exec(p.FirstName, p.LastName, p.ID)
	if nil != err {
		return
	}

	id, err = rs.RowsAffected()
	if nil != err {
		return
	}

	return
}

//DeletePerson 删除
func (p *Person) DeletePerson() (id int64, err error) {
	rs, err := db.SQLDB.Exec("DELETE FROM person WHERE id=?", p.ID)
	if nil != err {
		return
	}
	id, err = rs.RowsAffected()
	if nil != err {
		return
	}
	return
}
