package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/souviks72/flipcart-api/db/sql"
	hs "github.com/souviks72/flipcart-api/handlers"
	pr "github.com/souviks72/flipcart-api/params"
)

func main() {
	e := echo.New()

	dbConn, err := sql.NewDatabase()
	if err != nil {
		fmt.Printf("Err connecting to db: %+v\n", err)
	}
	err = dbConn.Client.AutoMigrate(
		&pr.Product{},
		&pr.CartItem{},
		&pr.Company{},
		&pr.Category{})
	if err != nil {
		e.Logger.Fatalf("Error migrating db: %+v\n", err)
		return
	}

	svc := hs.InitAPIService(dbConn)

	e.POST("/api/v1/products", svc.GetAllProducts)
	e.GET("/api/v1/products/:id", svc.GetProductById)

	e.POST("/api/v1/cart", svc.AddItemToCart)
	e.GET("/api/v1/cart", svc.GetAllCartItems)
	e.PATCH("/api/v1/cart", svc.UpdateCart)
	e.DELETE("/api/v1/cart/:id", svc.DeleteItemFromCart)

	e.Logger.Fatal(e.Start(":8000"))
}
