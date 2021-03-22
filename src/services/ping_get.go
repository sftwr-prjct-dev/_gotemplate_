package services

import (
	"context"
	"net/http"

	"gitlab.com/coinprofile/services/gotemplate/src/config"
	"gitlab.com/coinprofile/services/gotemplate/src/utils"
)

type PingReq struct{}

type PingRes struct {
	Version string `json:"version,omitempty"`
}

func (d *PingReq) Validate() (isValid bool, errs []error) {
	return len(errs) == 0, errs
}

// Controller returns the result of the logic
func (d *PingReq) Controller(ctx context.Context, cfg *config.Config) (status int, msg string, data interface{}, err error) {

	data = &PingRes{
		Version: cfg.Version,
	}

	return http.StatusOK, "RequestCompleted", data, err
}

func (d *PingReq) GetParamsMap() utils.QueryMap {
	return nil
}

func (d *PingReq) New() utils.RequestData {
	instance := PingReq{}
	return &instance
}
