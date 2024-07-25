package main

// start a server
// create all restful routes -> get, post, delete

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type people struct {
	id   string `json:"id"`
	name string `json:"name"`
	age  int    `json:"age"`
}

var peopleList = []people{
	{id: "1", name: "John", age: 35},
	{id: "2", name: "Doe", age: 30},
	{id: "3", name: "Jane", age: 25},
}

func main() {
	server(":9000")
}

func server(connection string) {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/people", getAllPeople)
	router.GET("/people/:id", getPersonById)
	router.POST("/people", addPerson)
	router.DELETE("/people/delete/:id", deletePerson)

	// testing query parameters
	router.GET("/hello/:name", printOnServer)
	router.GET("/hello/:name/*action", printWithAnAction)
	// ?name=John&age=35
	router.GET("/hello", printWithQuery)

	router.Run(connection)
}

func printOnServer(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"message": "Hello " + name})
}

func printWithAnAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{"message": "Hello " + name + " you are " + action})
}

func printWithQuery(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	age := c.Query("age")

	c.JSON(http.StatusOK, gin.H{"message": "Hello " + name + " you are " + age + " years old"})
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the home page"})
}

func getAllPeople(c *gin.Context) {
	// c.IndentedJSON(http.StatusOK, peopleList)
	c.JSON(http.StatusOK, peopleList)
}

func getPersonById(c *gin.Context) {
	id := c.Param("id")
	for _, person := range peopleList {
		if person.id == id {
			c.JSON(http.StatusOK, person)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

func addPerson(c *gin.Context) {
	var person people

	// throw error if request is invalid
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	peopleList = append(peopleList, person)
	c.JSON(http.StatusCreated, person)
}

func deletePerson(c *gin.Context) {
	id := c.Param("id")

	for i, person := range peopleList {
		if person.id == id {
			peopleList = append(peopleList[:i], peopleList[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "person deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "person not found"})
}
