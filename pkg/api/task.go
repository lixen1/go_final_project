package api

import (
	"encoding/json"
	"go_final_project/pkg/db"
	"net/http"
	"time"
)

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" || task.Date == now.Format(DateFormat) {
		task.Date = now.Format(DateFormat)
		return nil
	}

	t, err := time.Parse(DateFormat, task.Date)
	if err != nil {
		return err
	}

	var next string
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(DateFormat)
		} else {
			task.Date = next
		}
	}
	return nil
}

func writeJSON(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}

	return nil
}

func writeError(w http.ResponseWriter, errorMessage string, status int) {
	errorResponse := map[string]string{"error": errorMessage}

	w.Header().Set("Content-Type", "application/json charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorResponse)
}

type emptyStruct struct{}

var EmptyResponse = emptyStruct{}

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodPut:
		updateTaskHandler(w, r)
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodDelete:
		deleteTaskHandler(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
