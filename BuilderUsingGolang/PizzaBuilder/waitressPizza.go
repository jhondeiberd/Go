package PizzaBuilder


type waitressPizza struct {
	builderPizza iBuilderPizza
}

func NewWaitressPizza(b iBuilderPizza) *waitressPizza {
	return &waitressPizza{
		builderPizza: b,
	}
}

func (d *waitressPizza) SetBuilder(b iBuilderPizza) {
	d.builderPizza = b
}

func (d *waitressPizza) BuildPizza() pizza {
	d.builderPizza.setDough()
	d.builderPizza.setSauce()
	d.builderPizza.setTopping()
	return d.builderPizza.getPizza()
}
