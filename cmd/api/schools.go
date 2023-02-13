//Filename cmd/api/schools.go
package main

import(
	"net/http"
	"fmt"
)

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Created a new Schoool...")	
}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "School displayed...")	
}
