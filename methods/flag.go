package main

import (
	"flag"
	"fmt"
)

func sleep() {
	// var period = flag.Duration("sleep", 1*time.Second, "sleep period")
	// flag.Parse()
	// fmt.Printf("Sleep for %s\n", period)
	// time.Sleep(*period)
	// fmt.Println("finished")

	var period = celsiusFlag("celsius", Celsius(25), "sleep period")
	flag.Parse()
	fmt.Printf("Sleep for %s\n", period)
	// time.Sleep(*period)
	// fmt.Println("finished")
}

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%f °C", c)
}

func (c *Celsius) Set(s string) error {
	var end string
	var val float64

	fmt.Sscanf(s, "%f%s", &val, &end)

	switch end {
	case "°C":
		*c = Celsius(val)

		return nil
	}
	return fmt.Errorf("Celsius Parse Error")
}

func celsiusFlag(name string, val Celsius, desc string) *Celsius {

	flag.CommandLine.Var(&val, name, desc)
	return &val
}


