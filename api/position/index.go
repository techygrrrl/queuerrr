package position

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

	query := r.URL.Query()
	userId := query.Get("user_id")
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(api_utils.ErrorJson("missing query param: user_id"))
		return
	}

	db, err := api_utils.NewDatabaseClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}
	repo := api_utils.NewQueueRepository(db)

	position := repo.GetPosition(userId)
	payload := map[string]int{
		"position": position,
	}
	res, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	w.Write(res)
}
