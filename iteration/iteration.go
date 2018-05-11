package iteration

//Repeat concatenates the inputString on itself for repeatCount number of times.
func Repeat(inputString string, timesToRepeat int) (outputString string) {
	for i := 0; i < timesToRepeat; i++ {
		outputString += inputString
	}
	return;
}