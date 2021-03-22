package handlers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/coinprofile/services/gotemplate/src/config"
	"gitlab.com/coinprofile/services/gotemplate/src/utils"

	log "github.com/sirupsen/logrus"
)

func Handler(cfg *config.Config, _reqData utils.RequestData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData := _reqData.New()

		if r.Method == http.MethodGet || r.Method == http.MethodDelete {
			log.Debug("Processing request params")
			queryParams := reqData.GetParamsMap()
			if queryParams != nil {
				for k := range queryParams {
					queryParams[k] = r.URL.Query().Get(k)
				}

				jsonBody, err := json.Marshal(queryParams)

				if err != nil {
					utils.JSONResponse(w, http.StatusBadRequest, &utils.Response{Success: false, Message: "utils.DecodingErrorMsg", Data: nil})
					return
				}
				if err := json.Unmarshal(jsonBody, &reqData); err != nil {
					log.WithError(err).Error("utils.DecodingErrorMsg")
					utils.JSONResponse(w, http.StatusBadRequest, &utils.Response{Success: false, Message: "utils.DecodingErrorMsg", Data: nil})
					return
				}
			}
		} else {
			log.Debug("Processing request body")
			if err := utils.DecodeJSONBody(w, r, &reqData); err != nil {
				log.WithError(err).Error("utils.DecodingErrorMsg")
				utils.JSONResponse(w, http.StatusBadRequest, &utils.Response{Success: false, Message: "utils.DecodingErrorMsg", Data: nil})
				return
			}
		}

		log.Debug("Processed request param or body")

		if isValid, errs := reqData.Validate(); !isValid {
			utils.JSONResponse(w, http.StatusBadRequest, &utils.Response{Success: false, Message: "utils.InvalidReqDataErrorMsg", Data: utils.ResponseErrors{Errors: utils.CreateErrorResponse(errs...)}})
			return
		}

		log.Debug("Processing valid request")

		status, msg, data, err := reqData.Controller(r.Context(), cfg)
		if err != nil {
			utils.JSONResponse(w, status, &utils.Response{Success: false, Message: msg, Data: utils.ResponseErrors{Errors: utils.CreateErrorResponse(err)}})
			return
		}

		if status == http.StatusNoContent {
			utils.JSONResponse(w, status, nil)
			return
		}

		utils.JSONResponse(w, status, &utils.Response{Success: true, Message: msg, Data: data})
	}
}
