package PizzaBuilder

type hawaiianBuilderPizza struct {
	dough string
	sauce   string
	topping      string
}

func newhawaiianBuilderPizza() *hawaiianBuilderPizza {
	return &hawaiianBuilderPizza{}
}

func (b *hawaiianBuilderPizza) setDough() {
	b.dough = "hawaiian dough"
}

func (b *hawaiianBuilderPizza) setSauce() {
	b.sauce = "hawaiian sauce"
}

func (b *hawaiianBuilderPizza) setTopping() {
	b.topping = "hawaiian topping"
}

func (b *hawaiianBuilderPizza) getPizza() pizza {
	return pizza{
		Dough:   b.dough,
		Sauce: b.sauce,
		Topping:      b.topping,
	}
}
