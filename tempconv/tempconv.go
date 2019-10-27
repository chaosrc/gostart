package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

const (
	AbsoluteZero Celsius = -273.15
	FreezingC    Celsius = 0
	BollingC     Celsius = 100
)

// func main() {
// 	var f Celsius = 100
// 	fmt.Printf("%s\n", f)
// }
