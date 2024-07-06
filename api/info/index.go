package info

import (
	"encoding/json"
	"net/http"

	api_utils "github.com/techygrrrl/queuerrr/api_utils"
)

func Json(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := api_utils.NewDatabaseClient()
	if err != nil {
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	repo := api_utils.NewQueueRepository(db)

	users, err := repo.GetAll()
	if err != nil {
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	payload := map[string]any{
		"total": len(users),
		"users": users,
	}

	res, err := json.Marshal(payload)
	if err != nil {
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	w.Write(res)
}
