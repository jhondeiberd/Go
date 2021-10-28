package abstractFactory

type ISandwich interface {
	GetTyp() string
	GetBread()	string
	GetSalade() string
	GetTopping() string
}


type sandwich struct {
	typ string
	bread string
	salade string
	topping string
}
func (s *sandwich) GetTyp() string {
	return s.typ
}
func (s *sandwich) GetBread() string {
	return s.bread
}
func (s *sandwich) GetSalade() string {
	return s.salade
}
func (s *sandwich) GetTopping() string {
	return s.topping
}

