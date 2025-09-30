package api

import (
	"go_final_project/pkg/db"
	"net/http"
)

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		writeError(w, "error get tasks list:"+err.Error(), http.StatusNotFound)
		return
	}

	resp := TasksResp{
		Tasks: tasks,
	}

	if err := writeJSON(w, resp); err != nil {
		writeError(w, "JSON encoding error"+err.Error(), http.StatusInternalServerError)
		return
	}
}
