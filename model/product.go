package model

type Product struct {
	Id       int    `json:"product_id"`
	Name     string `json:"product_name"`
	Category *int   `json:"category"`
	Owner    int    `json:"product_owner"`
	Price    int    `json:"product_price"`
	Stock    *int   `json:"stock"`
}

type ProductJSON struct {
	Id       int       `json:"product_id"`
	Name     string    `json:"product_name"`
	Category *Category `json:"category"`
	Owner    *Owner    `json:"product_owner"`
	Price    int       `json:"product_price"`
	Stock    *int      `json:"stock"`
}

type Products []*ProductJSON

func NewProduct(id int, name string, category *int, owner int, price int, stock *int) *Product {
	return &Product{
		Id:       id,
		Name:     name,
		Category: category,
		Owner:    owner,
		Price:    price,
		Stock:    stock,
	}
}

func NewProductJSON(product *Product, category *Category, owner *Owner) *ProductJSON {
	return &ProductJSON{
		Id:       product.Id,
		Name:     product.Name,
		Category: category,
		Owner:    owner,
		Price:    product.Price,
		Stock:    product.Stock,
	}
}

func (pr *Products) Length() int {
	return len(*pr)
}

func (pr *Product) GetName() string {
	return pr.Name
}
