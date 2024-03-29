package httpapi

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/tedkimdev/go-web-skeleton/api/openapi"
)

func RespondWithError(w http.ResponseWriter, code int, err error) {
	if err == nil {
		slog.Info("do not call respondWithError with a nil err")
		return
	}

	slog.Error("response with error", slog.Any("error", err))
	RespondWithJSON(w, code, openapi.Error{
		Code:    int32(code),
		Message: err.Error(),
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if payload != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			slog.Error("failed to marshal json", slog.Any("error", err))
			w.WriteHeader(http.StatusInternalServerError)
			response, _ := json.Marshal(openapi.Error{
				Code:    http.StatusInternalServerError,
				Message: "failed to marshal json",
			})
			_, _ = w.Write(response)
			return
		}
		w.WriteHeader(code)
		_, _ = w.Write(response)
	}
}
