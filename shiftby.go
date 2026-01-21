// You can edit this code!
// Click here and start typing.
package sprint


// You can edit this code!
// Click here and start typing.

func ShiftBy(r rune, step int) rune {
	a := int(r)

	if a >= 'a' && a <= 'z' {
		step = step % 26
		shifted := a + step

		if shifted > 'z' {
			shifted = 'a' + (shifted - 'z' - 1)
		}

		return rune(shifted)
	}

	return r
}



