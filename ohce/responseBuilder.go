package main

type ResponseBuilder struct {
	name string
}

type ResponseBuilderI interface {
	BuildNameResponse(input string, clock ClockI) string
	BuildNormalResponse(input string) string
}

func (rb *ResponseBuilder) BuildNameResponse(input string, clock ClockI) string {
	rb.name = input
	switch {
	case clock.GetTime() < 12 && clock.GetTime() >= 6:
		return "¡Good morning " + input + "!\n"
	case clock.GetTime() < 18 && clock.GetTime() >= 12:
		return "¡Good afternoon " + input + "!\n"
	case clock.GetTime() < 21 && clock.GetTime() >= 18:
		return "¡Good evening " + input + "!\n"
	case clock.GetTime() < 6 || clock.GetTime() >= 21:
		return "¡Good night " + input + "!\n"
	default:
		return ""
	}
}

func (rb *ResponseBuilder) BuildNormalResponse(input string) string {
	if isStopWord(input) {
		return "Goodbye " + rb.name
	}
	if isPalindrome(input) {
		return input + "\n¡Nice palindrome!"
	}
	return reverseString(input) + ""
}

func isStopWord(s string) bool {
	if s == "Stop!" {
		return true
	}
	return false
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
