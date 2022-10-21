package ejemplo_2

import "strings"

type Payment struct {
	Amount float64
	Type   string
	Status string
	Detail string
}

func (p *Payment) IsValid() bool {
	return p.Amount > 0
}

func (p *Payment) SanitizeType() {
	p.Type = strings.TrimSpace(p.Type)
}
