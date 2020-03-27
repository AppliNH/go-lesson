package _types

type Metric float64

const (
	Meter      Metric = 1.
	Kilometer  Metric = Meter * 1000.
	Centimeter Metric = Meter / 100.

	Feet Metric = Centimeter * 30.48
	Inch Metric = Centimeter * 2.54
)
