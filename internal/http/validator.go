package http

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func BindAndValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		fmt.Println("error binding", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(i); err != nil {
		fmt.Println("error validating", err.Error())
		return err
	}

	return nil
}
