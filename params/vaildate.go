package params

import (
	"errors"
	"fmt"
)

func (p *ProductRequestParams) Ok(d Database) (TransformedProductSearchParams, error) {
	tr := TransformedProductSearchParams{}
	tr.Size = p.Size
	tr.ProductName = p.Attributes.Name
	tr.MinPrice = p.Attributes.MinPrice
	tr.MaxPrice = p.Attributes.MaxPrice

	if p.Attributes.Company != "" {
		company, err := d.GetCompanyByName(p.Attributes.Company)
		if err != nil {
			fmt.Printf("Product filter validation failed: %+v\n", err)
			return tr, err
		}
		if company.ID == 0 {
			return tr, errors.New("company does not exist")
		}
		tr.CompanyID = company.ID
	}

	if p.Attributes.Category != "" {
		category, err := d.GetCategoryByName(p.Attributes.Category)
		if err != nil {
			fmt.Printf("Product filter validation failed: %+v\n", err)
			return tr, err
		}
		if category.ID == 0 {
			return tr, errors.New("category does not exist")
		}
		tr.CategoryID = category.ID
	}

	return tr, nil
}
