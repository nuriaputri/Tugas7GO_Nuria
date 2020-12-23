package controllers

import (
	"Tugas7NuriaPutri/echo-rest/models"
	"net/http"

	"github.com/labstack/echo"
)

//FetchAllSuppliers ...
func FetchSuppliers(c echo.Context) (err error) {

	result, err := models.FetchSuppliers()

	return c.JSON(http.StatusOK, result)
}

func AddSupplier(c echo.Context) (err error) {
	result, err := models.AddSupplier(c)

	return c.JSON(result.Status, result)
}

func UpdateSupplier(c echo.Context) (err error) {
	result, err := models.UpdateSupplier(c)

	return c.JSON(result.Status, result)
}

func DeleteSupplier(c echo.Context) (err error) {
	result, err := models.DeleteSupplier(c)

	return c.JSON(result.Status, result)
}
