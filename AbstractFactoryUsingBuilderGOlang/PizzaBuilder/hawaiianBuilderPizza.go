package PizzaBuilder

type HawaiianBuilderPizza struct {
	dough   string
	sauce   string
	topping string
}

func newhawaiianBuilderPizza() *HawaiianBuilderPizza {
	return &HawaiianBuilderPizza{}
}

func (b *HawaiianBuilderPizza) setDough() {
	b.dough = "hawaiian dough"
}

func (b *HawaiianBuilderPizza) setSauce() {
	b.sauce = "hawaiian sauce"
}

func (b *HawaiianBuilderPizza) setTopping() {
	b.topping = "hawaiian topping"
}

func (b *HawaiianBuilderPizza) getPizza() Pizza {
	return Pizza{
		Dough:   b.dough,
		Sauce:   b.sauce,
		Topping: b.topping,
	}
}
