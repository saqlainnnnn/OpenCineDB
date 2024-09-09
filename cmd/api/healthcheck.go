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
		app.serverErrorResponse(w, r, err)
	}
}