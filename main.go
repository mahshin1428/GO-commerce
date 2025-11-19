package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type products struct{
		ID    	int  `json:"id"` //[`json:"id"` convert ID-->id]    // captital name for public and small name for private access
		Title   string  `json:"title"`
		Description string  `json:"description"`
		Price    string    `json:"price"`
}

var productsList []products

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, " Tui amre chinbi na ami tor baper khaloto bhai er chachato bhai er sumondir polar chacha")
}



func createProducts(w http.ResponseWriter, r *http.Request){

	corsHandler(w)
	

	 if r.Method != "POST"{
		http.Error(w, "please give me POST request", 400)  //200--> okay , 201-->created something, 400--> bad information, 404-->request not found in the server, 500-->internal server error 
		return
	}

	var newProduct products
	decoder := json.NewDecoder(r.Body)   //NewDecoder json thake struct e convert kore backend thake frontend e pathay
	err:= decoder.Decode(&newProduct)
	if err != nil{
		fmt.Println(err)
		http.Error(w, "please give a valid json", 400)
		return
	}

	
	newProduct.ID =len(productsList) + 1   // assigning id to the product, as User cannot assign id to the products
	productsList = append(productsList, newProduct)


	encoder := json.NewEncoder(w)     //Encoder json thake struct e convert kore backend thake frontend e pathay  
	encoder.Encode(newProduct)

}	

func getProducts(w http.ResponseWriter, r *http.Request){
	corsHandler(w)

	preFlightRq(w, r)

	 if r.Method != "GET"{
		http.Error(w, "please give me GET request", 400)  //200--> okay , 201-->created something, 400--> bad information, 404-->request not found in the server, 500-->internal server error 
		return
	}
	encoder := json.NewEncoder(w)     //Encoder json thake struct e convert kore backend thake frontend e pathay  
	encoder.Encode(productsList)

}

func corsHandler(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-allow-Headers", "Content-Type")
	w.Header().Set("Content-Type","application/json")
}

func preFlightRq(w http.ResponseWriter, r *http.Request){
	if r.Method =="OPTIONS"{        //browser send a preflight request to check is user allowed to send complex request or not
		w.WriteHeader(200)
	}
}



func main() {

	mux := http.NewServeMux() //router

	mux.HandleFunc("/hello", helloHandler)

	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/products", getProducts)

	mux.HandleFunc("/create-products",createProducts)

	fmt.Println("Server running on: 3000")

	err := http.ListenAndServe(":3000",  mux)

	if err != nil {
		fmt.Println("Error starting the server")
	}
}

func init(){
	p1 := products{
		ID: 1 ,
		Title: "Banana",
		Description: "this is a food with yellow colour",
		Price: "20 taka/kg",

	}
	p2 := products{
		ID: 2 ,
		Title: "Mango",
		Description: "this is a food with yellowishRed colour",
		Price: "80 taka/kg",

	}
	p3 := products{
		ID: 3 ,
		Title: "Oragne",
		Description: "this is a food with organe colour",
		Price: "250 taka/kg",

	}
	p4 := products{
		ID: 4 ,
		Title: "Lichi",
		Description: "this is a food with Red colour",
		Price: "300 taka/kg",

	}

	productsList = append(productsList, p1)
	productsList = append(productsList, p2)
	productsList = append(productsList, p3)
	productsList = append(productsList, p4)
}

