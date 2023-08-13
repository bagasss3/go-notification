package constant

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// http errors
var (
	ErrInvalidArgument     = echo.NewHTTPError(http.StatusBadRequest, "invalid argument")
	ErrNotFound            = echo.NewHTTPError(http.StatusNotFound, "record not found")
	ErrInternal            = echo.NewHTTPError(http.StatusInternalServerError, "internal system error")
	ErrUnauthenticated     = echo.NewHTTPError(http.StatusUnauthorized, "unauthenticated")
	ErrUnauthorized        = echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	ErrRefreshTokenExpired = echo.NewHTTPError(http.StatusNotAcceptable, "refresh token expired")
	ErrPasswordMismatch    = echo.NewHTTPError(http.StatusBadRequest, "password mismatch")
	ErrAlreadyExists       = echo.NewHTTPError(http.StatusBadRequest, "record already exists")
	ErrFieldEmpty          = echo.NewHTTPError(http.StatusBadRequest, "requirement field empty")
)

// httpValidationOrInternalErr return valdiation or internal error
func HttpValidationOrInternalErr(err error) error {
	switch t := err.(type) {
	case validator.ValidationErrors:
		_ = t
		errVal := err.(validator.ValidationErrors)

		fields := map[string]interface{}{}
		for _, ve := range errVal {
			fields[ve.Field()] = fmt.Sprintf("Failed on the '%s' tag", ve.Tag())
		}

		return echo.NewHTTPError(http.StatusBadRequest, dump(fields))
	default:
		return ErrInternal
	}
}

// Dump to json using json marshal
func dump(i interface{}) string {
	return string(toByte(i))
}

// ToByte :nodoc:
func toByte(i interface{}) []byte {
	bt, _ := json.Marshal(i)
	return bt
}
