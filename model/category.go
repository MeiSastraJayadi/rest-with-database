package model

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Categories []*Category

func NewCategory(id int, name string) *Category {
	return &Category{
		Id:   id,
		Name: name,
	}
}

func (c *Categories) Length() int {
	return len(*c)
}

func (c *Category) GetName() string {
	return c.Name
}
