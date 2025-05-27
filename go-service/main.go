package main

import (
        "net/http"

        "github.com/gin-gonic/gin"
)

type Person struct {
        ID   string `json:"id"`
        Name string `json:"name"`
        Age  int    `json:"age"`
}

var people = []Person{
        {ID: "1", Name: "Alice", Age: 30},
        {ID: "2", Name: "Bob", Age: 25},
}

func main() {
        r := gin.Default()

        r.GET("/persons", getPersons)
        r.GET("/persons/:id", getPersonByID)
        r.POST("/persons", createPerson)

        r.Run(":9000") // Listen and serve on localhost:9000
}

func getPersons(c *gin.Context) {
        c.JSON(http.StatusOK, people)
}

func getPersonByID(c *gin.Context) {
        id := c.Param("id")
        for _, p := range people {
                if p.ID == id {
                        c.JSON(http.StatusOK, p)
                        return
                }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
}

func createPerson(c *gin.Context) {
        var newPerson Person
        if err := c.ShouldBindJSON(&newPerson); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
        }
        people = append(people, newPerson)
        c.JSON(http.StatusCreated, newPerson)
}
