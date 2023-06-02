package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Customer struct {
	Nama string `validate:"required"`
	Email string `validate:"required,email"`
	Alamat string `validate:"required"`
	Umur int `validate:"gte=17,lte=35"`
}

func TestVariableValidation(c echo.Context) error  {
	v := validator.New()

	email := "dwi@gmail.com"

	err := v.Var(email, "required,email")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Email not valid",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})
}

func TestStructValidation(c echo.Context) error {
	v := validator.New()

	cust := Customer{
		Nama: "Dwi",
		Email: "dwi@gmail.cm",
		Alamat: "Dieng",
		Umur: 22,
	}

	err := v.Struct(cust)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})
}