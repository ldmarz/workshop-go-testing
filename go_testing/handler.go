package go_testing

import (
	"net/http"
	"strconv"
)

func SomeFancyHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sendErrorArg, ok := r.URL.Query()["send_error"]
		if !ok || len(sendErrorArg) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("{\"error\": \"400\",\"message\": \"send_error argument is required\"}"))
			return
		}

		sendError, _ := strconv.ParseBool(sendErrorArg[0])

		if sendError {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("{\"error\": \"400\",\"message\": \"You ask me for and error\"}"))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{\"name\": \"charmeleon\",\"type\": \"fire\"}"))

	}
}
