package routes

import (
	"Tugas7NuriaPutri/echo-rest/controllers"

	"github.com/labstack/echo"
)

//SuppliersRoute ...
func SuppliersRoute(g *echo.Group) {

	g.POST("/list", controllers.FetchSuppliers)

	g.POST("/add", controllers.AddSupplier)

	g.POST("/update", controllers.UpdateSupplier)

	g.POST("/delete", controllers.DeleteSupplier)

}
