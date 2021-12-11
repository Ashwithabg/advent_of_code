package models

type Octopus struct {
	EnergyLevel int
	IsFlashing  bool
}

func (o *Octopus) Energize() {
	o.EnergyLevel++
}

func (o *Octopus) Flash() {
	o.IsFlashing = true
}

func (o *Octopus) Reset() {
	o.EnergyLevel = 0
}
