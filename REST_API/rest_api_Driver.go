package REST_API

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CarCompany struct { // https://appdividend.com/2022/06/22/golang-serialize-json-string/
	ID          string `json:"id"`
	Name        string `json:"name"`
	Founder     string `json:"founder"`
	Date        int    `json:"date"`
	PatchNumber int    `json:"patch_number"`
}

var carBrands = []CarCompany{
	{ID: "1", Name: "Ferrari", Founder: "Enzo Ferrari", Date: 1939, PatchNumber: 10},
	{ID: "2", Name: "Mercedes", Founder: "Karl Benz", Date: 1926, PatchNumber: 10},
	{ID: "3", Name: "Ford", Founder: "Henry Ford", Date: 1903, PatchNumber: 10},
	{ID: "4", Name: "Lamborghini", Founder: "Ferruccio Lamborghini", Date: 1963, PatchNumber: 10},
	{ID: "5", Name: "Toyota", Founder: "Kiichiro Toyoda", Date: 1957, PatchNumber: 10},
}

func POSTCars(ctx *gin.Context) { // CREATE
	var newCarBrand CarCompany
	if err := ctx.BindJSON(&newCarBrand); err != nil {
		return // return ctx https://stackoverflow.com/questions/45239409/empty-return-in-func-with-return-value-in-golang
	}
	carBrands = append(carBrands, newCarBrand)
	ctx.IndentedJSON(http.StatusCreated, newCarBrand)
}

func GETCars(ctx *gin.Context) { // READ
	ctx.IndentedJSON(http.StatusOK, carBrands)
}

func GETCarByID(id string) (*CarCompany, error) {
	for i, b := range carBrands {
		if b.ID == id {
			return &carBrands[i], nil
		}
	}
	return nil, errors.New("Car company not found")
}
func PATCHcars(ctx *gin.Context) {
	id, check := ctx.GetQuery("id")
	if !check {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing ID"})
		return
	}

	// return JSON with certain id of car brand
	// =========================================================
	var err error
	var car *CarCompany
	for i, c := range carBrands {
		if c.ID == id {
			car, err = &carBrands[i], nil
			break
		} else {
			car, err = nil, errors.New("")
		}

	}
	// =========================================================
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Car company not found!"})
		return
	}
	car.PatchNumber = car.PatchNumber + 1

	ctx.IndentedJSON(http.StatusOK, car)
}

func DELETEcars(ctx *gin.Context) {
	// Go is a garbage collected language.
	// You are not supposed to, and you cannot delete objects from memory.
	// It is the garbage collector's duty and responsibility to do so, and it does this automatically.
	// The garbage collector will properly remove objects from memory when they become unreachable.
	// https://stackoverflow.com/questions/42066797/how-to-delete-struct-object-in-go
	// -----------------------------------------------------------------------------------------------
	// delete from map by ID ---> remove(map, ID) if ID exists
}
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
