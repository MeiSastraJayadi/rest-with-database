package model

type Owner struct {
	Id      int    `json:"owner_id"`
	Name    string `json:"owner_name"`
	Address string `json:"owner_address"`
	Phone   string `json:"phone_number"`
}

type Owners []*Owner

func NewOwner(id int, name string, address string, phone string) *Owner {
	return &Owner{
		Id:      id,
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

func (ow *Owners) Length() int {
	return len(*ow)
}

func (ow *Owner) GetName() string {
	return ow.Name
}
