package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/golang/gddo/httputil/header"
)

// JSONResponse sends a json http response
func JSONResponse(w http.ResponseWriter, statusCode int, data *Response) {
	w.WriteHeader(statusCode)
	log.Debug("Header written")
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.WithError(err).Error("JSONEncodingError")
			JSONResponse(w, http.StatusInternalServerError, &Response{false, "ServerErrorMsg", ResponseErrors{Errors: []string{"InternalServerError"}}})
			return
		}
		log.WithFields(log.Fields{
			"success":     data.Success,
			"responseMsg": data.Message,
			"statusCode":  statusCode,
		}).Debug("RequestCompleted")
	} else {
		log.WithFields(log.Fields{
			"statusCode": statusCode,
		}).Debug("RequestCompleted")
	}

}

// DecodeJSONBody decodes a JSON request
func DecodeJSONBody(w http.ResponseWriter, r *http.Request, data interface{}) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			return errors.New("InvalidHeaderValueError")
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	decodedData := json.NewDecoder(r.Body)
	decodedData.DisallowUnknownFields()

	err := decodedData.Decode(&data)

	if err != nil && err != io.EOF {
		return err
	}

	return nil

}
