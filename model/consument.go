package model

type Consument struct {
	Id    int    `json:"consument_id"`
	Name  string `json:"consument_name"`
	Email string `json:"email"`
	Phone string `json:"phone_number"`
}

type Consuments []*Consument

func NewConsument(id int, name string, email string, phone string) *Consument {
	return &Consument{
		Id:    id,
		Name:  name,
		Email: email,
		Phone: phone,
	}
}

func (cs *Consuments) Length() int {
	return len(*cs)
}

func (c *Consument) GetName() string {
	return c.Name
}
