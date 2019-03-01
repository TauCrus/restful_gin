package apis

import (
	"fmt"
	"log"
	"net/http"
	"restful_gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IndexAPI 首页
func IndexAPI(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

// AddPersonAPI 新增的接口
func AddPersonAPI(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	person := models.Person{FirstName: firstName, LastName: lastName}
	raRows, err := person.AddPerson()
	if nil != err {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("insert successful %d", raRows)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

//GetPersonsAPI 查询所有的接口
func GetPersonsAPI(c *gin.Context) {
	person := models.Person{}
	persons, err := person.GetPersons()
	if nil != err {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}

//GetPersonAPI 查询单条记录的接口
func GetPersonAPI(c *gin.Context) {
	id := c.Param("id")
	p1 := models.Person{}
	p2, err := p1.GetPerson(id)
	if nil != err {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"person": p2,
	})
}

//UpdatePersonAPI 更新接口
func UpdatePersonAPI(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if nil != err {
		log.Fatalln(err)
	}
	person := models.Person{ID: id}
	err = c.Bind(&person)
	if nil != err {
		log.Fatalln(err)
	}

	raRows, err := person.UpdatePerson()

	msg := fmt.Sprintf("Update person %d successful %d", person.ID, raRows)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

//DeletePersonAPI 删除接口
func DeletePersonAPI(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if nil != err {
		log.Fatalln(err)
	}

	person := models.Person{ID: id}
	raRows, err := person.DeletePerson()

	msg := fmt.Sprintf("Delete person %d successful %d", id, raRows)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
