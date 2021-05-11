package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/wusphinx/gin-swagger-toy/docs"
)

// @Summary Get toy
// @Description get toy by ID
// @Accept  json
// @Produce  json
// @Param   toyID     path    int     true        "path参数：玩具ID"
// @Param   pageNum   query   int     true        "query参数：页码"
// @Param   pageSize   query   int     false        "每页数量"
// @Param   requestID     header  string     true        "header参数：request id"
// @Success 200 {string} string	"ok"
// @Router /v1/toys/{toyID} [get]
func GetToy(c *gin.Context) {}

// @Summary Create toy
// @Description create toy
// @Accept  json
// @Produce  json
// @Param   data     body    types.ToyReq     true        "新建玩具body入参"
// @Success 200 {string} string	"ok"
// @Router /v1/toys/ [post]
func CreateToy(c *gin.Context) {}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/v1/toys/:toyID", GetToy)
	r.POST("/v1/toys/", CreateToy)

	r.Run()
}
