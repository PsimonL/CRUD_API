package REST_API_MONGO

import "go.mongodb.org/mongo-driver/bson/primitive"

type CarCompany struct { // https://appdividend.com/2022/06/22/golang-serialize-json-string/
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Founder string             `json:"founder" bson:"founder"`
	Date    int                `json:"date" bson:"date"`
}

//
//type CarCompany struct { // https://appdividend.com/2022/06/22/golang-serialize-json-string/
//	ID      primitive.ObjectID `json:"id" bson:"_id"`
//	Name    string             `json:"name" bson:"name"`
//	Founder string             `json:"founder" bson:"founder"`
//	Date    int                `json:"date" bson:"date"`
//}

//var carBrands = []CarCompany{
//	{ID: "1", Name: "Ferrari", Founder: "Enzo Ferrari", Date: 1939, PatchNumber: 10},
//	{ID: "2", Name: "Mercedes", Founder: "Karl Benz", Date: 1926, PatchNumber: 10},
//	{ID: "3", Name: "Ford", Founder: "Henry Ford", Date: 1903, PatchNumber: 10},
//	{ID: "4", Name: "Lamborghini", Founder: "Ferruccio Lamborghini", Date: 1963, PatchNumber: 10},
//	{ID: "5", Name: "Toyota", Founder: "Kiichiro Toyoda", Date: 1957, PatchNumber: 10},
//}
