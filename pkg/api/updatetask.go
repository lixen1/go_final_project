package api

import (
	"encoding/json"
	"go_final_project/pkg/db"
	"net/http"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeError(w, "invalid JSON format: "+err.Error(), http.StatusBadRequest)
		return
	}

	if task.ID == "" {
		writeError(w, "id is empty", http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeError(w, "title is empty", http.StatusBadRequest)
		return
	}

	if err := checkDate(&task); err != nil {
		writeError(w, "invalid date: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		writeError(w, "error while updating task: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := writeJSON(w, EmptyResponse); err != nil {
		writeError(w, "error writing empty response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
