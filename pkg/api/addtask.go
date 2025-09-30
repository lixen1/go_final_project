package api

import (
	"encoding/json"
	"go_final_project/pkg/db"
	"net/http"
	"strconv"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeError(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeError(w, "Title cannot be empty", http.StatusBadRequest)
		return
	}

	if err := checkDate(&task); err != nil {
		writeError(w, "Invalid date: "+err.Error(), http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeError(w, "DB error: "+err.Error(), http.StatusServiceUnavailable)
		return
	}

	task.ID = strconv.Itoa(int(id))

	idResponse := map[string]string{"id": task.ID}

	if err = writeJSON(w, idResponse); err != nil {
		writeError(w, "JSON encoding error"+err.Error(), http.StatusServiceUnavailable)
		return
	}

}
