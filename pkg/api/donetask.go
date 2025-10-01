package api

import (
	"go_final_project/pkg/db"
	"net/http"
	"time"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, "id can't be empty", http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, "error get task"+err.Error(), http.StatusBadRequest)
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeError(w, "error delete task:"+err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		nextDate, err := NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			writeError(w, "error calculating next date: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = db.UpdateDate(nextDate, id)
		if err != nil {
			writeError(w, "error update task date:"+err.Error(), http.StatusBadRequest)
			return
		}
	}

	if err := writeJSON(w, EmptyResponse); err != nil {
		writeError(w, "error writing empty response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
