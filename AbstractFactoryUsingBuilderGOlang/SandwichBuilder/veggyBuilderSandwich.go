package SandwichBuilder

type VeggyBuilderSandwich struct {
	bread   string
	salade  string
	topping string
}

func newVeggyBuilderSandwich() *VeggyBuilderSandwich {
	return &VeggyBuilderSandwich{}
}

func (b *VeggyBuilderSandwich) setBread() {
	b.bread = "veggy bread"
}

func (b *VeggyBuilderSandwich) setSalade() {
	b.salade = "veggy salade"
}

func (b *VeggyBuilderSandwich) setTopping() {
	b.topping = "veggy topping"
}

func (b VeggyBuilderSandwich) getSandwich() Sandwich {
	return Sandwich{
		Bread:   b.bread,
		Salade:  b.salade,
		Topping: b.topping,
	}
}
