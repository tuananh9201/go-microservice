package main

type Item struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ItemCreate struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (ItemCreate) TableName() string {
	return "items"
}

type ItemUpdate struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ItemFilter struct {
	Name string `json:"name"`
}
