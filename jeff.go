package main

type Jeff struct {
  Ops int
  Active bool
}

func (j Jeff) FillPrimary(gb *GrainBag) {
  gb.FillPrimary()
  return
}

func (j Jeff) FillSecondary(gb *GrainBag) {
  gb.FillSecondary()
  return
}

func (j Jeff) PutInPlastic(gb *GrainBag) {
  gb.PutInPlastic()
  return
}
