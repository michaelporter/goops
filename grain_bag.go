package main

type GrainBag struct {
	primary   bool
	secondary bool
	plastic   bool
}

func (gb *GrainBag) FillPrimary() {
	gb.primary = true
}

func (gb *GrainBag) FillSecondary() {
	gb.secondary = true
}

func (gb *GrainBag) PutInPlastic() {
	gb.plastic = true
}
