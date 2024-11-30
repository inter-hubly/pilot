package valueobject

type Phone struct {
	id     string
	number string
}

func NewPhone(id, number string) Phone {
	return Phone{
		id:     id,
		number: number,
	}
}

func (p *Phone) Id() string {
	return p.id
}

func (p *Phone) Number() string {
	return p.number
}
