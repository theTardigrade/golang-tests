package tests

import "fmt"

func Message(expectedOutput, actualOutput interface{}) (msg string, fail bool) {
	if actualOutput != expectedOutput {
		msg = fmt.Sprintf("Test failed:\n\nEXPECTED:\n\t%v\nBUT GOT:\n\t%v", expectedOutput, actualOutput)
		fail = true
	}

	return
}
