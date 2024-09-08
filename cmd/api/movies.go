package main

import (
	"fmt"
	"net/http"
	"time"

	"greelight.alexedwards.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a movie")
}

func (app*application) showMovieHandler(w http.ResponseWriter, r *http.Request)  {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
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
		http.Error(w, "the server encountered a problem couldnt process your request", http.StatusInternalServerError)
	}
}