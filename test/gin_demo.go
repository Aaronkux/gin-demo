package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type User2 struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      int    `json:"age" binding:"required,validAge"`
}

func validAge(fl validator.FieldLevel) bool {
	age := fl.Field().Int()
	return age > 0 && age < 150
}

func middle() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before")
		c.Next()
		fmt.Println("after")
	}
}

func demo() {
	// core.Viper()
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validAge", validAge)
	}

	r.POST("/test", func(c *gin.Context) {
		var user User2
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  err.Error(),
			})
			return
		}
		fmt.Printf("%T", user.Age)
		c.JSON(200, gin.H{
			"msg":  "ok",
			"code": 0,
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		if file, err := c.FormFile("file"); err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  err.Error(),
			})
			return
		} else {
			err := c.SaveUploadedFile(file, "./"+file.Filename)
			if err != nil {
				c.JSON(200, gin.H{
					"code": 1,
					"msg":  err.Error(),
				})
			} else {
				c.JSON(200, gin.H{
					"code": 0,
					"msg":  "ok",
				})
			}
		}
	})

	v1 := r.Group("v1").Use(middle())
	v1.GET("testGroup", func(c *gin.Context) {
		fmt.Println("testGroup")
		c.JSON(200, gin.H{
			"msg":  "ok",
			"code": 0,
		})
	})

	r.Run(":8080")
}
