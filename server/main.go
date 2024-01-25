package main;


import (
	"net/http"
	"example.com/database"
	
)

func main() {

	
	database.Connect();

	http.ListenAndServe(":8080",nil);

}
