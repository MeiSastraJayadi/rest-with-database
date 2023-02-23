package model

type Transaction struct {
  Id int `json:"transaction_id"` 
  Buyer int `json:"buyer_id"`
  Product int `json:"product_id"`
  Total int `json:"total_pcs"`
}

type Transactions []*Transaction

func NewTransaction(id int, buyer int, product int, pcs int) *Transaction {
  return &Transaction{
    Id: id, 
    Buyer: buyer,
    Product: product,
    Total: pcs,
  }
}

func (trs *Transactions) Length() int {
  return len(*trs)
}

func (tr *Transaction) GetName() string {
  return ""
}



