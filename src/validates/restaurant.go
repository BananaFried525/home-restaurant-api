package validates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTableList(c *gin.Context) (GetTableListRequest, error) {
	var request GetTableListRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return request, err
	}

	return request, nil
}

func CreateTableOrder(c *gin.Context) (CreateTableOrderRequest, error) {
	var request CreateTableOrderRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return request, err
	}

	return request, nil
}
