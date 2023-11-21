package params

type Database interface {
	GetAllProducts() ([]Product, error)
	GetProductById(int) (Product, error)
	GetAllCartItems() ([]CartItem, error)
	SaveCartItem(CartItem) error
	GetCartItemByProductId(int) (CartItem, error)
	DeleteItemFromCart(int) error
	GetCartItemById(int) (CartItem, error)
	GetCompanyById(int) (Company, error)
	GetCompanyByName(string) (Company, error)
	GetCategoryById(int) (Category, error)
	GetCategoryByName(string) (Category, error)
}

type Product struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Price       int
	Description string
	Stock       int
	ImageFolder string
	CompanyID   int
	Company     Company
	CategoryID  int
	Category    Category
}

type Company struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type CartItem struct {
	ID        int `gorm:"primaryKey"`
	ProductID int
	Product   Product
	Quantity  int
}
