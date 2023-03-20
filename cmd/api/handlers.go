// Filename cmd/api/schools.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/LianniMatthews/orange/internal/data"
)

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	//struct to hold school provided by request
	var input struct {
		Name    string   `json:"name"`
		Level   string   `json:"level"`
		Contact string   `json:"contact"`
		Phone   string   `json:"phone"`
		Email   string   `json:"email"`
		Website string   `json:"website,omitempty"`
		Address string   `json:"address"`
		Mode    []string `json:"mode"`
	}
	//new json.Decoder instance
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	//Print request
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	//fmt.Fprintf(w, "show details of %d\n", id)
	school := data.School{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "University of Belmopan",
		Level:     "University",
		Contact:   "Lianni Matthews",
		Phone:     "323-4545",
		Website:   "https://uob.edu.bz",
		Address:   "17 Apple Avenue",
		Mode:      []string{"Blended", "Online", "Face-to-Face"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)

	if err != nil {
		app.logger.Println(err)
		app.serveErrorResponse(w, r, err)
		return
	}

}
