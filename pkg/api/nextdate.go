package api

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func afterNow(date, now time.Time) bool {
	return date.After(now)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {

	if now.IsZero() {
		return "", errors.New("now is zero date in NextDate func")
	}
	if dstart == "" {
		return "", errors.New("dstart in NextDate func is empty")
	}

	if repeat == "" {
		return "", errors.New("repeat in NextDate func is empty")
	}

	dstartTime, err := time.Parse(DateFormat, dstart)
	if err != nil {
		log.Println("dstart doesn't match the format")
		return "", err
	}

	repeatSplit := strings.Split(repeat, " ")

	if repeatSplit[0] == "d" {
		repeatNum, err := strconv.Atoi(repeatSplit[len(repeatSplit)-1])

		if err != nil || repeatNum < 0 || repeatNum > 400 {
			return "", err
		}

		for {
			dstartTime = dstartTime.AddDate(0, 0, repeatNum)
			if afterNow(dstartTime, now) {
				break
			}

		}

	} else if repeatSplit[0] == "y" {

		for {
			dstartTime = dstartTime.AddDate(1, 0, 0)
			if afterNow(dstartTime, now) {
				break
			}

		}
	} else {
		return "", errors.New("unavailable repeat format")
	}

	return dstartTime.Format(DateFormat), nil
}

func nextDayHandler(w http.ResponseWriter, r *http.Request) {

	nowParam := r.URL.Query().Get("now")
	dateParam := r.URL.Query().Get("date")
	repeatParam := r.URL.Query().Get("repeat")

	var nowTime time.Time

	if nowParam == "" {
		nowTime = time.Now()

	} else {
		var err error

		nowTime, err = time.Parse(DateFormat, nowParam)
		if err != nil {
			http.Error(w, "Parse error!", http.StatusBadRequest)
			return
		}
	}

	nextDate, err := NextDate(nowTime, dateParam, repeatParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := io.WriteString(w, nextDate); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
