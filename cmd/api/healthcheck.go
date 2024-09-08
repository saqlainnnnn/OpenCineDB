package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//created envelope map for the response 
	env := envelope{
		"status": "availiable",
		"system_info": map[string]string{
			"environment" : app.config.env,
			"version" : version,
		},
	}
	//writing json response
	err := app.writeJson(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "the server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}