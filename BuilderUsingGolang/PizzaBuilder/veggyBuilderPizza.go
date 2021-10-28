package PizzaBuilder

type veggyBuilderPizza struct {
	dough string
	sauce   string
	topping     string
}

func newVeggyBuilderPizza() *veggyBuilderPizza {
	return &veggyBuilderPizza{}
}

func (b *veggyBuilderPizza) setDough() {
	b.dough = "veggy dough"
}

func (b *veggyBuilderPizza) setSauce() {
	b.sauce = "veggy sauce"
}

func (b *veggyBuilderPizza) setTopping() {
	b.topping = "veggy topping"
}

func (b *veggyBuilderPizza) getPizza() pizza {
	return pizza{
		Dough:   b.dough,
		Sauce: b.sauce,
		Topping: b.topping,
	}
}
