package middlewares

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"gitlab.com/coinprofile/services/gotemplate/src/utils"
)

func ResponseHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				log.WithField("Panic message", err).Error("Recovered from an handler error")
				utils.JSONResponse(w, http.StatusInternalServerError, &utils.Response{Success: false, Message: "services.InternalServerError", Data: nil})
			}
		}()
		next.ServeHTTP(w, r)
	})
}
