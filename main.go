package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Users = []User{}

func main() {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", GetUsers)
		userRoutes.POST("/", CreateUser)
		userRoutes.PUT("/:id", UpdateUser)    // PUT /users/{id}
		userRoutes.DELETE("/:id", DeleteUser) // DELETE /users/{id}
	}

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}

}

func GetUsers(c *gin.Context) {
	c.JSON(200, Users) //[]
}

func CreateUser(c *gin.Context) {
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	reqBody.ID = uuid.New().String()

	Users = append(Users, reqBody)
	c.JSON(200, gin.H{
		"error": false,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	for i, u := range Users {
		if u.ID == id {
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error":   true,
		"message": "Invalid user id",
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, u := range Users {
		if u.ID == id {
			// t := [1,40,50,60,75]
			// t [:2] == [1,40]
			//t [2+1:] == [60,75]
			// [1,40,60,75]
			Users = append(Users[:i], Users[i+1:]...)

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}
