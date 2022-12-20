package main

import (
	"awesomeProject4/REST_API"
	"fmt"
	//"encoding/json"
)

func main() {
	// https://github.com/gin-gonic/gin/blob/master/README.md
	var i int
	choice_map := make(map[int]string)
	choice_map[1] = "Basic CRUD API"
	choice_map[2] = "Mongo DB CRUD API"
	fmt.Printf("%s - 1\n", choice_map[1])
	fmt.Printf("%s - 2\n", choice_map[2])
	fmt.Print("Your choice: ")
	//fmt.Scan(&i)
	i = 2
	if i == 1 {
		fmt.Println("You picked 1st option:", choice_map[i])
		fmt.Println()
		REST_API.DriverCode()
	} else if i == 2 {
		fmt.Println("You picked 2nd option:", choice_map[i])
		fmt.Println()
		//REST_API_MONGO.Driver_Code_MONGO()
		//// https://gorm.io/docs/create.html
	} else {
		fmt.Println("There are only 2 options to chose from, 1 or 2.")
	}
}
