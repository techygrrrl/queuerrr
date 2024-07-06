package next

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

	user, err := repo.NextInQueue()
	if err != nil {
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	res, err := json.Marshal(user)
	if err != nil {
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	w.Write(res)
}
