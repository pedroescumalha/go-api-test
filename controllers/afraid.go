package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AfraidOf struct {
	IsHeAfraid bool
}

func isChuckNorrisAfraid(subject string) bool {
	chuckNorrisIsAfraidOf := []string{
		"himself",
		"chucknorris",
		"ChuckNorris",
		"Chucknorris",
		"chuckNorris",
		"chuck-norris",
	}

	for i := 0; i < len((chuckNorrisIsAfraidOf)); i++ {
		if chuckNorrisIsAfraidOf[i] == subject {
			return true
		}
	}

	return false
}

func GetIsChuckNorrisAfraid(context *gin.Context) {
	subject := context.Param("subject")

	context.IndentedJSON(
		http.StatusOK,
		AfraidOf{
			IsHeAfraid: isChuckNorrisAfraid(subject),
		},
	)
}
