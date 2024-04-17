/*
Package towfer is a module to implement imaginary bakery that has
a holiday offer where you can buy two cookies for the price of one.
*/
package twofer

// ShareWith is a function which share a bread with a given name.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
