package ejemplo_1

type Payment struct {
	Type   string
	Amount float64
	Status string
	Detail string
}

func (p *Payment) IsValid() bool {
	return p.Amount > 0
}
