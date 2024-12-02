package nextdate

import (
	"net/http"
	"time"
)

func HandlerNextDate(w http.ResponseWriter, r *http.Request) {

	now := r.URL.Query().Get("now")
	date := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

	nowPars, err := time.Parse(DateFormat, now)
	if err != nil {
		http.Error(w, "ошибка парсинга даты", http.StatusBadRequest)
		return
	}

	nextDate, err := NextDate(nowPars, date, repeat)
	if err != nil {
		http.Error(w, "ошибка расчета даты", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}
