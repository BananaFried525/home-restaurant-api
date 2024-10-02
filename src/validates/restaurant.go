package validates

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTable(c *gin.Context) (*GetTableRequest, error) {
	var request GetTableRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return nil, err
	}

	log.Println(&request)
	return &request, nil
}

func AddTable(c *gin.Context) (*AddTableRequest, error) {
	var request AddTableRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return nil, err
	}

	return &request, nil
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

func GetTableDetail(c *gin.Context) (*GetTableDetailRequest, error) {
	var request GetTableDetailRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return nil, err
	}

	log.Println(&request)
	return &request, nil
}

func OrderFood(c *gin.Context) (*OrderFoodRequest, error) {
	var request OrderFoodRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return nil, err
	}

	return &request, nil
}
