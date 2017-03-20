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

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(releases); err != nil {
		panic(err)
	}
}

func ReleaseCreate(w http.ResponseWriter, r *http.Request) {
	var release Release
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &release); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    new_r := release.Create()

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(new_r); err != nil {
        panic(err)
    }

}

func ReleaseShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var release Release
	result := release.Get(vars["releaseId"])

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}