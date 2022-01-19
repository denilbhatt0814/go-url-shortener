package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/denilbhatt0814/practice/go-urls/pkg/models"
	"github.com/denilbhatt0814/practice/go-urls/pkg/utils"
)


func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!!\n")
}

func PathHandler(w http.ResponseWriter, r *http.Request){
	// getting the path
	path := r.URL.Path

	// search the DB for path -> url KV
	if url,err := models.GetUrl(path); url!="" && err==nil {
		http.Redirect(w,r, url, http.StatusFound)
		return
	}

	fmt.Fprintf(w, "PAGE NOT FOUND!\n")
}


func AddPath(w http.ResponseWriter, r *http.Request){
	// enpty pathurl object
	pathurl := &models.PathUrl{}

	// parsing the req to the empty obj
	utils.ParseBody(r, pathurl)

	// storing the KV pairs in DB
	pu := pathurl.ShortenUrl()
	res, _ := json.Marshal(pu)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}	

