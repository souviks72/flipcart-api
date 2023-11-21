package params

type ProductRequestParams struct {
	Size       int                     `json:"size,omitempty"`
	Attributes ProductSearchAttributes `json:"attributes,omitempty"`
}

type ProductSearchAttributes struct {
	Name     string `json:"name,omitempty"`
	MinPrice int    `json:"min_price,omitempty"`
	MaxPrice int    `json:"max_price,omitempty"`
	Category string `json:"category,omitempty"`
	Company  string `json:"company,omitempty"`
}

type AddToCartRequestParams struct {
	Quantity  int `json:"quantity"`
	ProductId int `json:"product_id"`
}

type TransformedProductSearchParams struct {
	Size        int
	ProductName string
	MinPrice    int
	MaxPrice    int
	CategoryID  int
	CompanyID   int
}
