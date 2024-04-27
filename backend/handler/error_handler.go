package handler

import (
	"net/http"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/handler/presenter"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/logger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

const (
	statusInternalErr = http.StatusInternalServerError
	statusNotFound    = http.StatusNotFound
	statusBadRequest  = http.StatusBadRequest
	statusNoPermitted = http.StatusUnauthorized
	statusConflict    = http.StatusConflict

	msgInternalErr    = "SystemError has occurred"
	msgNotFoundErr    = "Resource NotFound"
	msgBadRequestErr  = "Request is invalid"
	msgNoPermittedErr = "No permission"
	msgConflictErr    = "Conflict has occurred"
)

type errResult struct {
	httpStatusCode int
	message        string
}

type ErrorHandler struct {
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

var errResultMap = map[error]*errResult{}

func init() {
	errResultMap[domain.ErrInternal] = newErrResult(statusInternalErr, msgInternalErr)
	errResultMap[domain.ErrNotFound] = newErrResult(statusNotFound, msgNotFoundErr)
	errResultMap[domain.ErrBadRequest] = newErrResult(statusBadRequest, msgBadRequestErr)
	errResultMap[domain.ErrParmission] = newErrResult(statusNoPermitted, msgNoPermittedErr)
	errResultMap[domain.ErrAlreadyExist] = newErrResult(statusConflict, msgConflictErr)
	errResultMap[domain.ErrConflict] = newErrResult(statusConflict, msgConflictErr)
}

func handleError(domainErr error, originErr error) error {
	if _, ok := originErr.(validator.ValidationErrors); ok {
		logger.ErrorWithParams("ValidationErrors has occurred", map[string]interface{}{
			"domainErr": domainErr.Error(),
			"originErr": originErr.Error(),
		})
	}

	if _, exist := domain.ErrMap[originErr]; exist {
		return originErr
	}

	logger.ErrorWithParams("SystemError has occurred", map[string]interface{}{
		"domainErr": domainErr.Error(),
		"originErr": originErr.Error(),
	})

	return domainErr
}

func (h *ErrorHandler) Handler(err error, c echo.Context) {
	e, exist := errResultMap[err]
	if !exist {
		e = errResultMap[domain.ErrInternal]
	}

	_ = presenter.BuildErrorResponse(
		c,
		e.httpStatusCode,
		e.message,
	)
}

func newErrResult(httpStatusCode int, message string) *errResult {
	return &errResult{
		httpStatusCode: httpStatusCode,
		message:        message,
	}
}
