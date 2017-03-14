package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	
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

func ReleaseShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	releaseId := vars["releaseId"]
	fmt.Fprintln(w, "Release show:", releaseId)
}