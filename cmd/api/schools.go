// Filename cmd/api/schools.go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LianniMatthews/orange/internal/data"
	"github.com/LianniMatthews/orange/internal/validator"
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
	//decode JSON request
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)

		return
	}

	//validate JSON input
	v := validator.New()

	//validation checks:
	v.Check(input.Name != "", "name", "must be provided")
	v.Check(len(input.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(input.Level != "", "level", "must be provided")
	v.Check(len(input.Level) <= 200, "level", "must not be more than 200 bytes long")

	v.Check(input.Contact != "", "contact", "must be provided")
	v.Check(len(input.Contact) <= 200, "contact", "must not be more than 200 bytes long")

	v.Check(input.Phone != "", "phone", "must be provided")
	v.Check(validator.Matches(input.Phone, validator.PhoneRX), "phone", "must be a valid phone number")

	v.Check(input.Email != "", "email", "must be provided")
	v.Check(validator.Matches(input.Email, validator.EmailRX), "email", "must be a valid email address")

	v.Check(input.Website != "", "website", "must be provided")
	v.Check(validator.ValidWebsite(input.Website), "website", "must be a valid URL")

	v.Check(input.Address != "", "address", "must be provided")
	v.Check(len(input.Address) <= 500, "address", "must not be more than 500 bytes long")

	v.Check(input.Mode != nil, "mode", "must be provided")
	v.Check(len(input.Mode) >= 1, "mode", "must contain at least 1 entry")
	v.Check(len(input.Mode) <= 5, "mode", "must contain at most 5 entries")
	v.Check(validator.Unique(input.Mode), "mode", "must not contain duplicate entries")

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
