// @title NHBook API
// @version 1.0
// @description This is the API for NHBook, a book management system.
// @host localhost:3030
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Api-Key
// @contact.name Nguyen Anh Quan

package main

import "github.com/NguyenAnhQuan-Dev/NKbook-API/internal/app"

func main() {
	app.Run()
}
