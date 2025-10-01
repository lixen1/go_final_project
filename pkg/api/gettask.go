package api

import (
	"go_final_project/pkg/db"
	"net/http"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, "id can't be empty", http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, "error get task:"+err.Error(), http.StatusBadRequest)
		return
	}

	if err := writeJSON(w, task); err != nil {
		writeError(w, "JSON encoding error"+err.Error(), http.StatusInternalServerError)
		return
	}
}
