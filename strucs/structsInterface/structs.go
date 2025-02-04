package structs

type Product struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Type Type `json:"type"`
	Price float32 `json:"price"`
	Tag []string `json:"tag"`
}

type Type struct {
	ID uint `json:"id"`
	Code string `json:"code"`
	Description string `json:"description"`
}


func (p *Product) TotalPrice() float32{
	return float32(p.Price) + 10.0
}


/* 
func (p *Product) SetName(name string) float32{
	p.Name = name
} */

func (p *Product) AddTag(tag ...string) {
	p.Tag = append(p.Tag, tag...)
}
