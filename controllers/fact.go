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

	var fact Fact = Fact{
		Message: resp.Value,
	}

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, "error fetching fact")
		return
	}

	context.IndentedJSON(http.StatusOK, fact)
}
