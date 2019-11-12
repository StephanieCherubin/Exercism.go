// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

const askBobAQuestion = "WHAT"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if askBobAQuestion == strings.ToUpper(askBobAQuestion) {
		return "Uppercase"
		// } else if condition {
		// 	return " "
		// } else if condition {
		// 	return " "
		// } else if condition {
		// 	return ""
	// } else {
	// 	return "Whatever"
	// }
}

/*
He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
If question === question.allCaps() && question.endsWith(?)
return "Calm down, I know what I'm doing!"

Bob answers 'Sure.' if you ask him a question, such as "How are you?".
if remarks[-1] == ? return "Sure."

He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
if remarks === remarks.allCaps()
	return "Whoa, chill out!"

If question === " ":
	return "Fine. Be that way!"
He says 'Fine. Be that way!' if you address him without actually saying anything.

else:
	return "Whatever"
He answers 'Whatever.' to anything else.
*/
