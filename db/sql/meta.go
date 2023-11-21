package sql

import (
	"fmt"

	pr "github.com/souviks72/flipcart-api/params"
)

func (d *Database) GetCompanyById(id int) (pr.Company, error) {
	var company pr.Company
	res := d.Client.First(&company, id)
	if res.Error != nil {
		fmt.Printf("Error fetching company by id: %+v\n", res.Error)
		return company, res.Error
	}

	return company, nil
}

func (d *Database) GetCompanyByName(name string) (pr.Company, error) {
	var company pr.Company
	res := d.Client.Where("name = ?", name).First(&company)
	if res.Error != nil {
		fmt.Printf("Error fetching company by id: %+v\n", res.Error)
		return company, res.Error
	}

	return company, nil
}

func (d *Database) GetCategoryById(id int) (pr.Category, error) {
	var category pr.Category
	res := d.Client.First(&category, id)
	if res.Error != nil {
		fmt.Printf("Error fetching company by id: %+v\n", res.Error)
		return category, res.Error
	}

	return category, nil
}

func (d *Database) GetCategoryByName(name string) (pr.Category, error) {
	var category pr.Category
	res := d.Client.Where("name = ?", name).First(&category)
	if res.Error != nil {
		fmt.Printf("Error fetching company by id: %+v\n", res.Error)
		return category, res.Error
	}

	return category, nil
}
