package controllers

import (
	"api/example/clients"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Fact struct {
	Message string `json:"fact"`
}

func GetFact(context *gin.Context) {
	resp, err := clients.GetChuckNorrisFact()

	if err != nil || resp.Value == "" {
		context.IndentedJSON(http.StatusInternalServerError, "error fetching fact")
		return
	}

	var fact Fact = Fact{
		Message: resp.Value,
	}

	context.IndentedJSON(http.StatusOK, fact)
}
