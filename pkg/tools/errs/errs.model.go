package errs

import "time"

type ApiErrorResponse struct {
	Status  int       `json:"status"`
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	Request string    `json:"request"`
	Detail  string    `json:"detail"`
}
