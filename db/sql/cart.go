package sql

import pr "github.com/souviks72/flipcart-api/params"

func (d *Database) GetAllCartItems() ([]pr.CartItem, error) {
	var cartItems []pr.CartItem
	res := d.Client.Preload("Product").Find(&cartItems)
	if res.Error != nil {
		return nil, res.Error
	}

	return cartItems, nil
}

func (d *Database) SaveCartItem(cartItem pr.CartItem) error {
	res := d.Client.Save(&cartItem)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *Database) GetCartItemByProductId(productId int) (pr.CartItem, error) {
	cartItem := pr.CartItem{
		ProductID: productId,
	}
	res := d.Client.Where("product_id = ?", productId).Preload("Product").Find(&cartItem)
	if res.Error != nil {
		return cartItem, res.Error
	}

	return cartItem, nil
}

func (d *Database) DeleteItemFromCart(cartItemId int) error {
	res := d.Client.Delete(&pr.CartItem{}, cartItemId)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *Database) GetCartItemById(cartItemId int) (pr.CartItem, error) {
	cartItem := pr.CartItem{}
	res := d.Client.First(&cartItem, cartItemId)
	if res.Error != nil {
		return cartItem, res.Error
	}

	return cartItem, nil
}
