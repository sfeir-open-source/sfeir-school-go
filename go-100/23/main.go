package main

type Voiture interface {
	MoveTo(x, y, z float64)
	Stop()
	IsMoving() bool
}
