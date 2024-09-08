package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//creating data which will hold the info of the json response
	data := map[string]string{
		"status": "availiable",
		"environment": app.config.env,
		"version": version,
	}
	//writing json response
	err := app.writeJson(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "the server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}