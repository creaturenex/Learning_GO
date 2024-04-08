package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func RepeatWithCount(char string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += char
	}
	return repeated
}
