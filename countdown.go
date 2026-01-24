
package sprint


func Countdown(n int) string {
	result := "" // declare result first

	for i := n; i > 0; i -= 2 {
		result += string(rune('0'+ i)) + ", "

	}

	result += "0!"
	return result // append the final zero

}