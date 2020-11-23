// Package twofer provides a function to share with names passed in as paremeters.
package twofer

import "fmt"

// ShareWith provides a response based on certain parameters
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
