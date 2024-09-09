package main

import (
	"fmt"
	"net/http"
	"time"

	"greelight.alexedwards.net/internal/data"
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