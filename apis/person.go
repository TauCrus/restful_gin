package apis

import (
	"net/http"
	"log"
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	. "restful_gin/models"
)

func IndexApi(c *gin.Context){
	c.String(http.StatusOK, "It works")
}
// 新增的接口
func AddPersonApi(c *gin.Context){
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	person := Person{FirstName:firstName, LastName: lastName}
	ra_rows, err:= person.AddPerson()
	if nil != err {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("insert successful %d", ra_rows)
	
	c.JSON(http.StatusOK, gin.H{
		"msg":msg,
	})
}
//查询所有的接口
func GetPersonsApi(c *gin.Context){
	person := Person{}
	persons,err := person.GetPersons()
	if nil != err {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}
//查询单条记录的接口
func GetPersonApi(c *gin.Context){
	id := c.Param("id")
	p1 := Person{}
	p2, err := p1.GetPerson(id)
	if nil != err {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"person": p2,
	})
}
//更新接口
func UpdatePersonApi(c *gin.Context){
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if nil != err {
		log.Fatalln(err)
	}
	person := Person{Id: id}
	err = c.Bind(&person)
	if nil != err {
		log.Fatalln(err)
	}

	ra_rows, err := person.UpdatePerson()
	
	msg := fmt.Sprintf("Update person %d successful %d", person.Id, ra_rows)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
//删除接口
func DeletePersonApi(c *gin.Context){
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if nil != err{
		log.Fatalln(err)
	} 

	person := Person{Id: id}
	ra_rows, err := person.DeletePerson()
	
	msg := fmt.Sprintf("Delete person %d successful %d", id, ra_rows)
	
	c.JSON(http.StatusOK, gin.H{
		"msg":msg,
	})
}