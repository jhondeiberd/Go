package SandwichBuilder

type hawaiianBuilderSandwich struct {
	bread string
	salade   string
	topping      string
}

func newhawaiianBuilderSandwich() *hawaiianBuilderSandwich {
	return &hawaiianBuilderSandwich{}
}

func (b *hawaiianBuilderSandwich) setBread() {
	b.bread = "hawaiian bread"
}

func (b *hawaiianBuilderSandwich) setSalade() {
	b.salade = "hawaiian salade"
}

func (b *hawaiianBuilderSandwich) setTopping() {
	b.topping = "hawaiian topping"
}

func (b *hawaiianBuilderSandwich) getSandwich() sandwich {
	return sandwich{
		Bread:   b.bread,
		Salade: b.salade,
		Topping:      b.topping,
	}
}
