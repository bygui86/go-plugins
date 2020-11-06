package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"plugin"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type Shipper interface {
	Name() string
	Currency() string
	CalculateRate(weight float32) float32
}

type config struct {
	shippingMethod string
	packageWeight  float32
}

func main() {
	cfg, cfgErr := loadConfig()
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}

	// 1. Search the plugins directory for a file with the same name as the shippingMethod provided through env-var
	plugin, plugErr := plugin.Open(fmt.Sprintf("plugins/%s.so", cfg.shippingMethod))
	if plugErr != nil {
		log.Fatal(plugErr)
	}

	// 2. Look for an exported symbol (function or variable)
	// as we expect that every plugin exports a single struct
	// that implements the Shipper interface with the name "Shipper"
	shipperSymbol, symbErr := plugin.Lookup("Shipper")
	if symbErr != nil {
		log.Fatal(symbErr)
	}

	// 3. Attempt to cast the plugin symbol to Shipper interface
	// this will allow us:
	// 	. to call the methods on the plugins if the plugin has implemented the required methods
	// 	. or fail if it does not implement them
	var shipper Shipper
	var ok bool
	shipper, ok = shipperSymbol.(Shipper)
	if !ok {
		log.Fatal("Invalid shipper type")
	}

	// 4. If everything is ok till now, we can proceed
	// calling the methods on our shipper interface object
	rate := shipper.CalculateRate(cfg.packageWeight)
	rate1Day := fmt.Sprintf("%.2f %s", rate, shipper.Currency())
	rate2Days := fmt.Sprintf("%.2f %s", rate-(rate*.20), shipper.Currency())
	rate7Days := fmt.Sprintf("%.2f %s", rate-(rate*.70), shipper.Currency())

	fmt.Println("")
	fmt.Println(shipper.Name())
	fmt.Println("")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Number of Days", "Rate"})
	table.Append([]string{"1 Day Express", rate1Day})
	table.Append([]string{"2 Days Shipping", rate2Days})
	table.Append([]string{"7 Days Shipping", rate7Days})
	table.Render()

	fmt.Println("")
}

func loadConfig() (*config, error) {
	method := os.Getenv("SHIPPING_METHOD")
	if method == "" {
		return nil, errors.New("shipping method not specified")
	}

	weight, err := strconv.ParseFloat(os.Getenv("PACKAGE_WEIGHT"), 32) // Load the plugin
	if err != nil {
		return nil, err
	}

	return &config{
		shippingMethod: method,
		packageWeight:  float32(weight),
	}, nil
}
