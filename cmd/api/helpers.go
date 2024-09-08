package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]interface{}


func (app *application) writeJson(w http.ResponseWriter, status int, data envelope, headers http.Header ) error {
	// marrshal encodes the data into json
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	// makes it easier to read 
	js = append(js, '\n')
	// the headers are added to the http response headers
	for key, value := range headers {
		w.Header()[key] = value
	}
	//adding json to the header, then writing status code and json response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}


func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10 , 64)

	if err != nil || id <1 { 
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}