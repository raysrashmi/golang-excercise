// Package twofer provides a function to generate a sharing message
// following the "one for X, one for me" pattern.
package twofer

import "fmt"

// ShareWith returns a string indicating sharing with the given name.
// If name is empty, it defaults to "you".  
func ShareWith(name string) string {
	if name=="" {
		name = "you"
	}
    return fmt.Sprintf("One for %s, one for me.", name)
}
