package data

import (
	"database/sql"
	"time"

	"greelight.alexedwards.net/internal/validator"
)

type Movie struct {
	ID			int64  		`json:"id"`
	CreatedAt 	time.Time	`json:"created_at"`
	Title 		string		`json:"title"`
	Year 		int32		`json:"year,omitempty"`
	Runtime 	Runtime		`json:"runtime,omitempty"`
	Genres 		[]string	`json:"genres,omitempty"`
	Version 	int32		`json:"version"`
}

type MovieModel struct {
	DB *sql.DB
}

func (m MovieModel) Insert(movie *Movie) error {
	return nil
}

func (m MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

func (m MovieModel) Update(movie *Movie) error {
	return nil
}

func (m MovieModel) Delete(movie *Movie) error {
	return nil
}

func ValidateMovie(v *validator.Validator, movie *Movie)  {
	
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title)<= 5000, "title", "must not be more than 5000 words" )
	
	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "year must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "year must be greater than 1888")

	v.Check(movie.Runtime !=0, "rutime", "must be provided")
	v.Check(movie.Runtime >0, "rutime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >=1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must be less than 5 genre")

	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}