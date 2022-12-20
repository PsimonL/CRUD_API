// ###############################################################################################################
// NOT FINISHED
// ###############################################################################################################
package REST_API_MONGO

import (	"fmt"
			"github.com/gin-gonic/gin"
			//"log"
			//"gopkg.in/mgo.v2/bson"
	)

// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/retrieve/
// https://gin-gonic.com/docs/
// https://www.makeuseof.com/golang-crud-api-mongodb-gin/
// https://codevoweb.com/crud-restful-api-server-with-golang-and-mongodb/

func rest_api_mongo_Driver(){
	fmt.Println("Working Mongo driver code")
	router := gin.Default()

	//var collection = ConnectDB()
	//router.POST("/post", func(c *gin.Context) {
	//	doc := bson.D{{"_id", 5555}, {"ad", "ag"}, {"ba", "bd"}, {"ce", 5555}}
	//	result, err := collection.InsertOne(context.TODO(), doc)
	//	print(err)
	//	print(result)
	//	c.IndentedJSON(http.StatusCreated, doc)
	//	fmt.Println("Isnerted: ", doc)
	//})
	router.POST("/post", CreateSingleDataSetPOST)
	router.Run("localhost:8080")
}
}