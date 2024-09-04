package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/medali/ramo-server/internal/app"
	"github.com/medali/ramo-server/internal/sources/openai"
)

func HelperController(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")
	
	var input app.HelperInput
	_ = json.NewDecoder(req.Body).Decode(&input)
	
	text := openai.HelperFunc(input)
	json.NewEncoder(res).Encode(text)
}



