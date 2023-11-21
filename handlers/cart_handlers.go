package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	pr "github.com/souviks72/flipcart-api/params"
)

func (s *APIService) GetAllCartItems(c echo.Context) error {
	return s.fetchCartItems(c)
}

func (s *APIService) AddItemToCart(c echo.Context) error {
	var reqParams pr.AddToCartRequestParams
	err := c.Bind(&reqParams)
	if err != nil {
		fmt.Printf("Error binding AddToCart params: %+v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"message": "Invalid Request Params",
		})
	}

	cartItem, err := s.Database.GetCartItemByProductId(reqParams.ProductId)
	if err != nil {
		fmt.Printf("Error fetching cart item by product id: %+v\n", cartItem)
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to add item to cart")
	}

	fmt.Printf("%+v\n", cartItem)

	if cartItem.Quantity != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Item already exists in cart, use Update API")
	}

	err = s.Database.SaveCartItem(cartItem)
	if err != nil {
		fmt.Printf("Error adding item to cart: %+v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to add item to cart")
	}

	return s.fetchCartItems(c)
}

func (s *APIService) UpdateCart(c echo.Context) error {
	var cartRequestParams pr.AddToCartRequestParams
	err := c.Bind(&cartRequestParams)
	if err != nil {
		fmt.Printf("Error binding cart update params: %+v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Request body must have product_id and quantity")
	}

	cartItem, err := s.Database.GetCartItemByProductId(cartRequestParams.ProductId)
	if err != nil {
		fmt.Printf("Error fetching cart item by product id: %+v\n", cartItem)
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to update cart")
	}

	if cartRequestParams.Quantity == 0 {
		err := s.Database.DeleteItemFromCart(cartItem.ID)
		if err != nil {
			fmt.Printf("Error deleting item from cart: %+v\n", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to remove item from cart")
		}
	}

	cartItem.Quantity = cartRequestParams.Quantity

	err = s.Database.SaveCartItem(cartItem)
	if err != nil {
		fmt.Printf("Error updating cart: %+v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to update cart")
	}

	return c.JSON(http.StatusOK, cartItem)
}

func (s *APIService) DeleteItemFromCart(c echo.Context) error {
	id := c.Param("id")
	cartItemId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Invalid id: %+v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}

	cartItem, err := s.Database.GetCartItemById(cartItemId)
	if err != nil {
		fmt.Printf("Error fetching cart item by product id: %+v\n", cartItem)
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to add item to cart")
	}

	if cartItem.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Cart Item with given id does not exist")
	}

	err = s.Database.DeleteItemFromCart(cartItemId)
	if err != nil {
		fmt.Printf("Error deleting cart item by id: %+v\n", cartItem)
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to delete item from cart")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Item deleted"})
}

func (s *APIService) fetchCartItems(c echo.Context) error {
	cartitems, err := s.Database.GetAllCartItems()
	if err != nil {
		fmt.Printf("Error Fetching Cart Items: %+v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to fetch cart items")
	}

	return c.JSON(http.StatusOK, cartitems)
}
