package main

//import "time"
//import "fmt"

func main() {
  jeff := Jeff{}
  gb := new(GrainBag)
  jeff.FillPrimary(gb)
  jeff.FillSecondary(gb)
  jeff.PutInPlastic(gb)
}

