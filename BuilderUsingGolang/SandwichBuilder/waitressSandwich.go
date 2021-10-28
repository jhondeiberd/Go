package SandwichBuilder

type waitressSandwich struct {
	builderSandwich iBuilderSandwich
}

func NewWaitressSandwich(b iBuilderSandwich) *waitressSandwich {
	return &waitressSandwich{
		builderSandwich: b,
	}
}

func (d *waitressSandwich) SetBuilder(b iBuilderSandwich) {
	d.builderSandwich = b
}

func (d *waitressSandwich) BuildSandwich() sandwich {
	d.builderSandwich.setBread()
	d.builderSandwich.setSalade()
	d.builderSandwich.setTopping()
	return d.builderSandwich.getSandwich()
}
