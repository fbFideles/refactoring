package pedido

import "refactoring/item"

func New(cpf string) *Pedido {
	return &Pedido{
		cpf: cpf,
	}
}

type Pedido struct {
	cpf   string
	cupom float32
	items []item.Item
}

func (p *Pedido) AddItem(name string, quantity int) {
	p.items = append(p.items, *item.New(name, quantity))
}

func (p *Pedido) AddCupom(porcentagem float32) {
	p.cupom = porcentagem
}
