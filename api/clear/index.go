package leave

import (
	"encoding/json"
	"net/http"

	"github.com/techygrrrl/queuerrr/api_utils"
)

func Json(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := api_utils.Authenticate(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	db, err := api_utils.NewDatabaseClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}
	repo := api_utils.NewQueueRepository(db)

	err = repo.ClearQueue()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	payload := map[string]string{
		"status": "successfully cleared the queue",
	}
	res, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	w.Write(res)
}
