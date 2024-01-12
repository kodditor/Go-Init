package main

import (
	"github.com/gin-gonic/gin"
)

type student struct {
	Name string `json: "name"`
	Age  int    `json: "age"`
}

var students = []student{
	{Name: "Kwabena", Age: 19},
	{Name: "Kwadwo", Age: 10},
}

func main() {
	router := gin.Default()

	router.GET("/students", getStudents)
	router.POST("/students", addStudent)
	router.GET("/students/:name", findStudent)
	router.DELETE("/students/:name", removeStudent)
	router.Run("localhost:8080")
}

func getStudents(c *gin.Context) {
	c.IndentedJSON(200, students)
}

func addStudent(c *gin.Context) {
	var newStudent student
	if err := c.BindJSON(&newStudent); err != nil {
		return
	}

	students = append(students, newStudent)
	c.IndentedJSON(201, newStudent)
}

func findStudent(c *gin.Context) {
	name := c.Param("name")

	for _, s := range students {
		if s.Name == name {
			c.IndentedJSON(200, s)
		}
	}
}

func removeStudent(c *gin.Context) {
	name := c.Param("name")
	for idx, s := range students {

		if s.Name == name {
			students = append(students[:idx], students[(idx+1):]...)
			c.IndentedJSON(204, students)
		}

	}
}
