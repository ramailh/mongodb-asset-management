package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramailh/mongodb-asset-management/models"
	"github.com/ramailh/mongodb-asset-management/services"
)

func FindBlog(c *gin.Context) {
	var param models.GetAll
	if err := c.Bind(&param); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": false, "message": err.Error()})
		return
	}

	data, err := services.Find(param)
	if err != nil {
		c.JSON(200, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "data": data})
}

func FindBlogByID(c *gin.Context) {
	var param models.MicroBlogging
	param.ID = c.Param("id")

	data, err := services.FindByID(param)
	if err != nil {
		c.JSON(200, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "data": data})
}

func InsertBlog(c *gin.Context) {
	var param models.MicroBlogging
	if err := c.BindJSON(&param); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": false, "message": err.Error()})
		return
	}

	data, err := services.Insert(param)
	if err != nil {
		c.JSON(200, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "data": data})
}

func UpdateBlog(c *gin.Context) {
	var param models.MicroBlogging
	if err := c.BindJSON(&param); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": false, "message": err.Error()})
		return
	}

	param.ID = c.Param("id")

	data, err := services.Update(param)
	if err != nil {
		c.JSON(200, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "data": data})
}

func DeleteBlog(c *gin.Context) {
	var param models.MicroBlogging
	param.ID = c.Param("id")

	data, err := services.Delete(param)
	if err != nil {
		c.JSON(200, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "data": data})
}
