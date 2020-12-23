package main

import (
	db "Tugas7NuriaPutri/echo-rest/db"
	routes "Tugas7NuriaPutri/echo-rest/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))

}
