package REST_API

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func DriverCode() {
	fmt.Println("Working REST_API")
	//// Resources:
	//// https://dev.to/devniklesh/crud-api-with-go-gin-framework-production-ready-52jd
	//// https://go.dev/doc/tutorial/web-service-gin
	//// https://www.youtube.com/watch?v=bj77B59nkTQ&t=920s
	//// https://circleci.com/blog/gin-gonic-testing/
	//// https://gin-gonic.com/docs/quickstart/
	//
	router := gin.Default() // handling different routes and end points of API
	//
	router.GET("/carBrands/getfunc", GETCars)
	//// curl http://localhost:8080/carBrands
	//
	router.POST("/carBrands/postfunc", POSTCars)
	//// curl http://localhost:8080/carBrand
	//
	router.PATCH("/carBrands/patchfunc", PATCHcars) // or PUT
	//// http://localhost:8080/carBrands/patchfunc?id=1
	//
	router.DELETE("/carBrands/deletefunc", DELETEcars)
	//// http://localhost:8080/carBrands/deletefunc?id=1
	//
	router.Run("localhost:8080")
	//
}
