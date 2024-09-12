package main

import (
	"fmt"
	"net/http"
	"time"

	"greelight.alexedwards.net/internal/data"
	"greelight.alexedwards.net/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title	string			`json:"title"`
		Year	int32			`json:"year"`
		Runtime	data.Runtime	`json:"runtime"`
		Genres	[]string		`json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	v.Check(input.Title != "", "title", "must be provided")
	v.Check(len(input.Title)<= 5000, "title", "must not be more than 5000 words" )

	v.Check(input.Year != 0, "year", "must be provided")
	v.Check(input.Year >= 1888, "year", "year must be greater than 1888")
	v.Check(input.Year <= int32(time.Now().Year()), "year", "year must be greater than 1888")

	v.Check(input.Runtime !=0, "rutime", "must be provided")
	v.Check(input.Runtime >0, "rutime", "must be a positive integer")

	v.Check(input.Genres != nil, "genres", "must be provided")
	v.Check(len(input.Genres) >=1, "genres", "must contain at least 1 genre")
	v.Check(len(input.Genres) <= 5, "genres", "must be less than 5 genre")

	v.Check(validator.Unique(input.Genres), "genres", "must not contain duplicate values")

	if !v.Valid() {
		app.failedValidResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app*application) showMovieHandler(w http.ResponseWriter, r *http.Request)  {
	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w,r)
		return
	}
	//created instance of movie struct
	movie := data.Movie {
		ID: id,
		CreatedAt: time.Now(),
		Title: "Casablanca",
		Runtime: 102,
		Genres: []string{"drama", "romance", "war"},
		Version: 1,
	}
	//this time encoding a struct
	err = app.writeJson(w, http.StatusOK, envelope{"movie" : movie}, nil)
	if err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(w,r,err)
	}
}