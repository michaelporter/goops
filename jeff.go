package main

type Jeff struct {
  Ops int
}

func (j *Jeff) FillPrimary(gb *GrainBag) {
  gb.FillPrimary()
  j.Ops = j.Ops + 1
  return
}

func (j *Jeff) FillSecondary(gb *GrainBag) {
  gb.FillSecondary()
  j.Ops = j.Ops + 1
  return
}

func (j *Jeff) PutInPlastic(gb *GrainBag) {
  gb.PutInPlastic()
  j.Ops = j.Ops + 1
  return
}
