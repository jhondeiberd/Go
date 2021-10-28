package PizzaBuilder

type VeggyBuilderPizza struct {
	dough   string
	sauce   string
	topping string
}

func newVeggyBuilderPizza() *VeggyBuilderPizza {
	return &VeggyBuilderPizza{}
}

func (b *VeggyBuilderPizza) setDough() {
	b.dough = "veggy dough"
}

func (b *VeggyBuilderPizza) setSauce() {
	b.sauce = "veggy sauce"
}

func (b *VeggyBuilderPizza) setTopping() {
	b.topping = "veggy topping"
}

func (b *VeggyBuilderPizza) getPizza() Pizza {
	return Pizza{
		Dough:   b.dough,
		Sauce:   b.sauce,
		Topping: b.topping,
	}
}
