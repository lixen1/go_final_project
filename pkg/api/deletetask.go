package api

import (
	"go_final_project/pkg/db"
	"net/http"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, "id can't be empty", http.StatusBadRequest)
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		writeError(w, "error task deletion: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := writeJSON(w, EmptyResponse); err != nil {
		writeError(w, "error writing empty response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
