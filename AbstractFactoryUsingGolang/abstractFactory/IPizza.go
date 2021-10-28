package abstractFactory

type IPizza interface {
	GetTyp() string
	GetDough() string
	GetSauce() string
	GetTopping() string
}


type pizza struct {
	typ     string
	dough   string
	sauce   string
	topping string
}
func (s *pizza) GetTyp() string {
	return s.typ
}
func (s *pizza) GetDough() string {
	return s.dough
}
func (s *pizza) GetSauce() string {
	return s.sauce
}
func (s *pizza) GetTopping() string {
	return s.topping
}

