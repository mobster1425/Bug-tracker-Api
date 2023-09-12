package main

import (
	"feyin/bug-tracker/config"
	_ "feyin/bug-tracker/docs"

	// "feyin/bug-tracker/initializer"
	"feyin/bug-tracker/routes"

	"github.com/gin-gonic/gin"

	// 	ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	//"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"
)

var DB *gorm.DB
var R = gin.Default()

func init() {
//	initializer.LoadEnv()
	var err error
	config.DB, err = config.DBconnect()
	if err != nil {
		panic(err)
	}

	//	R.LoadHTMLGlob("templates/*.html")
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	// url := ginSwagger.URL("http://localhost:8080/docs/doc.json")
	// R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// url := ginSwagger.URL("http://localhost:8080/docs/swagger.json")

	// url := ginSwagger.URL("http://localhost:8080/docs/swagger.json") // The url pointing to API definition
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	// R.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// routes.AdminRouts(R)
	// routes.UserRouts(R)
	routes.ProjectRoute(R)
	routes.UserRoutes(R)
	//R.Run()
	R.Run(":8080")
}
