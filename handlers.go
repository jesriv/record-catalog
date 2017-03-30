package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"io"
	
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func ReleaseIndex(w http.ResponseWriter, r *http.Request) {
	releases := GetReleases()

	jsonHeaders(w, http.StatusOK)
	jsonResponse(w, releases)
}

func ReleaseCreate(w http.ResponseWriter, r *http.Request) {
	var release Release

	body := readBody(r.Body)

	if err := json.Unmarshal(body, &release); err != nil {
		jsonHeaders(w, 422)
		jsonResponse(w, err)
	}

	new_release := release.Create()

	jsonHeaders(w, http.StatusCreated)
	jsonResponse(w, new_release)
}

func ReleaseShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var release Release
	result := release.Get(vars["releaseId"])

	jsonHeaders(w, http.StatusOK)
	jsonResponse(w, result)
}

func ReleaseUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var release Release

	body := readBody(r.Body)

	if err := json.Unmarshal(body, &release); err != nil {
		jsonHeaders(w, 422)
		jsonResponse(w, err)
	}

	updated_release := release.Update(vars["releaseId"])

	jsonHeaders(w, http.StatusOK)
	jsonResponse(w, updated_release)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var user User

	body := readBody(r.Body)

	if err := json.Unmarshal(body, &user); err != nil {
		jsonHeaders(w, 422)
		jsonResponse(w, err)
	}

	if result, err := user.Authenticate(); err != nil {
		jsonHeaders(w, 401)
		jsonResponse(w, err)		
	} else {
		jsonHeaders(w, http.StatusOK)
		jsonResponse(w, result)
	}
}

func jsonHeaders(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
}

func jsonResponse(w http.ResponseWriter, res interface{}) {
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func readBody(rBody io.ReadCloser) []byte {
	body, err := ioutil.ReadAll(io.LimitReader(rBody, 1048576))

	if err != nil {
		panic(err)
	}

	if err := rBody.Close(); err != nil {
		panic(err)
	}

	return body
}