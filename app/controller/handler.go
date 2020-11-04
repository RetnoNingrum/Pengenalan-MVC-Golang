package controller

import (
	"github.com/RetnoNingrum/Pengenalan-MVC-Golang/app/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddAntrianHandler(c *gin.Context){
	flag, err := model.addAtrian()
	if err != nil{
		c.JSON(http.StatusInternalServerError,map[string]interface{}{
			"status":"failed",
		})
		return
	}
	if flag {
		c.JSON(http.StatusOK,map[string]interface{}{
			"status":"success",
		})
	} else {
		c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status":"failed",
			"error":err,
		})
	}
}

func GetAntrianHandler(c *gin.Context){
	flag,err,resp := model.getAntrian()
	if err != nil{
		c.JSON(http.StatusInternalServerError,map[string]interface{}{
			"status":"failed",
			"massage":err.Error(),
		})
		return
	}
	if flag {
		c.JSON(http.StatusOK,map[string]interface{}{
			"status":"success",
			"data": resp,
		})
	} else {
		c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status":"failed",
			"massage":"unknown error",
		})
	}
}

func UpdateAntrianHandler(c *gin.Context){
	idAntrian := c.Param("idAntrian")
	err := model.updateAntrian(idAntrian)
	if err != nil{
		c.JSON(http.StatusInternalServerError,map[string]interface{}{
			"status":"failed",
			"massage":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,map[string]interface{}{
		"status":"success",
	})
}

func DeleteAntrianHandler(c *gin.Context){
	idAntrian := c.Param("idAntrian")
	err := model.deleteAntrian(idAntrian)
	if err != nil{
		c.JSON(http.StatusInternalServerError,map[string]interface{}{
			"status":"failed",
			"massage":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,map[string]interface{}{
		"status":"success",
	})
}

func PageAntrianHandler(c *gin.Context) {
	flag, err, result := model.getAntrian()
	var currentAntrian map[string]interface{}

	for _, item := range result {
		if item != nil {
			currentAntrian = item
			break
		}
	}

	if flag && len(result) > 0 {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"antrian": currentAntrian["id"],
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}