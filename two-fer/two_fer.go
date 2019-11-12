// https://golang.org/doc/effective_go.html#commentary
package twofer

/* Given a name, ShareWith function should return a string with the message "One for X, one for me"
However, if the name is missing, return the string:
"One for you, one for me"
*/
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
