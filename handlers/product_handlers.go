package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	pr "github.com/souviks72/flipcart-api/params"
)

func (s *APIService) GetAllProducts(c echo.Context) error {
	var params pr.ProductRequestParams
	err := c.Bind(&params)
	if err != nil {
		fmt.Printf("Error binding product search req body: %+v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	products, err := s.Database.GetAllProducts()
	if err != nil {
		fmt.Printf("Error fetching products: %+v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, products)
}

func (s *APIService) GetProductById(c echo.Context) error {
	id := c.Param("id")
	if len(id) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Id")
	}

	prdId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Id")
	}

	product, err := s.Database.GetProductById(prdId)
	if err != nil {
		fmt.Printf("Error fetching product by id: %+v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}
