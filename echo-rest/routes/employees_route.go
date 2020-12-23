package routes

import (
	"Tugas7NuriaPutri/echo-rest/controllers"

	"github.com/labstack/echo"
)

//CustomersRoute ...
func EmployeesRoute(g *echo.Group) {

	g.POST("/list", controllers.FetchAllEmployees)

	g.POST("/add", controllers.AddEmployees)

	//g.POST("/update", controllers.UpdateCustomer)

	//g.POST("/delete", controllers.DeleteCustomer)

}
