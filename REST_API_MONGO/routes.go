// ###############################################################################################################
// NOT FINISHED
// ###############################################################################################################

package REST_API_MONGO

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"time"
)

func CreateSingleDataSetPOST(ctx *gin.Context) {
	var collection = ConnectDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	post := new(CarCompany)
	defer cancel()
	if err := ctx.BindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	postPayload := CarCompany{
		ID:      primitive.NewObjectID(),
		Name:    post.Name,
		Founder: post.Founder,
		Date:    post.Date,
	}

	result, err := collection.InsertOne(ctx, postPayload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
}

func ReadWholeDataSetGET(ctx *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var collection = ConnectDB()
	id := ctx.Param("id")
	var result CarCompany
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(ctx, bson.M{"id": objID}).Decode(&result)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusCreated)
}
