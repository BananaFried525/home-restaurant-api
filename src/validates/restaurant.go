package validates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTable(c *gin.Context) *GetTableRequest {
	var request GetTableRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	return &request
}

func AddTable(c *gin.Context) *AddTableRequest {
	var request AddTableRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	return &request
}

func GetMenu(c *gin.Context) (*GetMenuRequest, error) {
	var request GetMenuRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "40001",
			"message": err.Error(),
		})
		return nil, err
	}

	return &request, nil
}

func AddOrder(c *gin.Context) (*AddOrderRequest, error) {
	var request AddOrderRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "40001",
			"message": err.Error(),
		})
		return nil, err
	}

	return &request, nil
}
