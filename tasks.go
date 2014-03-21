package main

func packPrimary(j *Jeff) {
	for {
		gb := <-readyForPrimary
		j.FillPrimary(gb)
		readyForSecondary <- gb
	}
}

func packSecondary(j *Jeff) {
	for {
		gb := <-readyForSecondary
		j.FillSecondary(gb)
		readyForPlastic <- gb
	}
}

func packInPlastic(j *Jeff) {
	for {
		gb := <-readyForPlastic
		j.PutInPlastic(gb)
		completed++
	}
}
