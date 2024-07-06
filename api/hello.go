package api

import (
	"encoding/json"
	"net/http"
	"time"

	api_utils "github.com/techygrrrl/queuerrr/api_utils"
)

func Json(w http.ResponseWriter, r *http.Request) {
	err := api_utils.Authenticate(r)
	if err != nil {
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	now := time.Now()
	bird := api_utils.Entry{
		CreatedAt: &now,
		Username:  "techygrrrl",
		UserId:    "176082690",
		Notes:     "This is a test entry",
	}
	res, err := json.Marshal(bird)
	if err != nil {
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	w.Write(res)
}
