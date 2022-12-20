package REST_API

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
