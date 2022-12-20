package REST_API

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
