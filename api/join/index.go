package join

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	api_utils "github.com/techygrrrl/queuerrr/api_utils"
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
	username := query.Get("username")
	if username == "" {
		w.Write(api_utils.ErrorJson("missing query param: username"))
		return
	}

	userId := query.Get("user_id")
	if userId == "" {
		w.Write(api_utils.ErrorJson("missing query param: user_id"))
		return
	}

	notes := query.Get("notes")
	notes, err = url.QueryUnescape(notes)
	if err != nil {
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	db, err := api_utils.NewDatabaseClient()
	if err != nil {
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	repo := api_utils.NewQueueRepository(db)
	err = repo.JoinQueue(userId, username, notes)
	if err != nil {
		fmt.Printf("here 1: %s", err.Error())

		var errorMessage string
		if strings.Contains(err.Error(), "duplicate key") {
			errorMessage = "user already in the queue"
		} else {
			errorMessage = err.Error()
		}

		w.Write(api_utils.ErrorJson(errorMessage))
		return
	}

	entry, err := repo.FindUser(userId)
	if err != nil {
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	res, err := json.Marshal(entry)
	if err != nil {
		w.Write(api_utils.ErrorJson(err.Error()))
		return
	}

	w.Write(res)
}
