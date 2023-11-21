package sql

import pr "github.com/souviks72/flipcart-api/params"

func (d *Database) GetAllProducts() ([]pr.Product, error) {
	var products []pr.Product
	res := d.Client.Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}

	return products, nil
}

func (d *Database) GetProductById(id int) (pr.Product, error) {
	var product pr.Product
	res := d.Client.First(&product, id)
	if res.Error != nil {
		return product, res.Error
	}

	return product, nil
}
