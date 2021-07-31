package controllers

type Posts struct {
	name            string
	link            string
	image           string
	shortDescrition string
	id              string
	londDescription string
	technologies    []string
}
type Post interface {
	New() Posts
	Delete() Posts
	Update() Posts
	Get() Posts
	GetAll() []Posts
}

func (p *Posts) New() Posts {
	p.name = "jwjewe"

	return *p
}
