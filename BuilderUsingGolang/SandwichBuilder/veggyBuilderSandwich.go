package SandwichBuilder

type veggyBuilderSandwich struct {
	bread string
	salade   string
	topping     string
}

func newVeggyBuilderSandwich() *veggyBuilderSandwich {
	return &veggyBuilderSandwich{}
}

func (b *veggyBuilderSandwich) setBread() {
	b.bread = "veggy bread"
}

func (b *veggyBuilderSandwich) setSalade() {
	b.salade = "veggy salade"
}

func (b *veggyBuilderSandwich) setTopping() {
	b.topping = "veggy topping"
}

func (b *veggyBuilderSandwich) getSandwich() sandwich {
	return sandwich{
		Bread:   b.bread,
		Salade: b.salade,
		Topping: b.topping,
	}
}
