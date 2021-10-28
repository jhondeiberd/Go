package SandwichBuilder

type HawaiianBuilderSandwich struct {
	bread   string
	salade  string
	topping string
}

func newhawaiianBuilderSandwich() *HawaiianBuilderSandwich {
	return &HawaiianBuilderSandwich{}
}

func (b *HawaiianBuilderSandwich) setBread() {
	b.bread = "hawaiian bread"
}

func (b *HawaiianBuilderSandwich) setSalade() {
	b.salade = "hawaiian salade"
}

func (b *HawaiianBuilderSandwich) setTopping() {
	b.topping = "hawaiian topping"
}

func (b HawaiianBuilderSandwich) getSandwich() Sandwich {
	return Sandwich{
		Bread:   b.bread,
		Salade:  b.salade,
		Topping: b.topping,
	}
}
