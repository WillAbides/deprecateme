package deprecateme

import (
	"fmt"
)

// HelloWorld returns a greeting
func HelloWorld(greetee string) string {
	if greetee == "" {
		greetee = "world"
	}
	return fmt.Sprintf("Hello, %s!", greetee)
}
