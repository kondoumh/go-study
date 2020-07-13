package Person

import "errors"

type Person struct {
	name string
}

func (p *Person) GetName() (string, error) {
	if p.name == "" {
		return "", errors.New("Name is not set")
	}
	return p.name, nil
}

func (p *Person) SetName(in_name string) {
	p.name = in_name
}
