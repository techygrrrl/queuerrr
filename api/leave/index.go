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
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	// Get the data
	query := r.URL.Query()
	userId := query.Get("user_id")
	if userId == "" {
		w.Write(api_utils.ErrorJson("missing query param: user_id"))
		return
	}

	db, err := api_utils.NewDatabaseClient()
	if err != nil {
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}
	repo := api_utils.NewQueueRepository(db)

	_, err = repo.FindUser(userId)
	if err != nil {
		w.Write(api_utils.ErrorJson("user not in queue"))
		return
	}

	err = repo.LeaveQueue(userId)
	if err != nil {
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	payload := map[string]string{
		"status": "successfully left queue",
	}
	res, err := json.Marshal(payload)
	if err != nil {
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	w.Write(res)
}
