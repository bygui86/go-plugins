package main

var Shipper shipper

type shipper struct{}

func (s shipper) Name() string {
	return "My delivery (MyDev)"
}

func (s shipper) Currency() string {
	return "EUR"
}

func (s shipper) CalculateRate(weight float32) float32 {
	cost := weight * 4.2
	tax := cost * .22
	return cost + tax
}
