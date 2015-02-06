//my name
//CSCI 325 Program 1 Part 2
//Check integer literal
//2015.02.04

//Decimal
//Octal
//Hex
//no underscores
//no binary
//optional l or L

//Copyright 2015, All rights reserved.

package main

//import "os"
import "strings"
import "strconv"
import "fmt"

func main() {
	teststrings := []string{"0", "0x", "0x_", " no ", "123", "1_23", " 2e", "1 2", "0xAB", "   1l", "  1  l 2"}
	for index, str := range teststrings {
		fmt.Println(teststrings[index], analyzeString(str))
	}
}

func analyzeString(s string) bool{
	//trim leading and trailing whitespace
	s = strings.TrimSpace(s)
	//lowercase everything for hex and 0x, etc
	s = strings.ToLower(s)
	//ignores underscores
	s = strings.Replace(s, "_", "", -1)
	//ignore a single leading l for type long
	if strings.HasSuffix(s, "l") {
		s = strings.TrimSuffix(s, "l")
	}
	
	switch{
		case strings.Contains(s, " "):
			//the trimmed string still contains a space
			return false
		case strings.HasPrefix(s, "0x"):
			//hex
			return checkValidHex(s)
		case strings.HasPrefix(s, "0"):
			//octal
			return checkValidOctal(s)
		default:
			//dec or garbage
			return checkValidDecimal(s)
	}
	return false
}

func checkValidHex(s string) bool {
	//define valid hex characters
	vHex := "1234567890abcde"
	//trim leading 0x from string
	s = strings.TrimPrefix(s, "0x")
	//if there was nothing after the "0x" then it isn't an integer
	if len(s) == 0 { return false }
	for _, char := range s {
		if !strings.ContainsAny(strconv.QuoteRune(char), vHex) {
			return false
		}
	}
	return true
}
func checkValidOctal(s string) bool {
	//define valid octal characters
	vOct := "01234567"
	//leading 0 can stay, it doesn't mess with letters
	//even if there's nothing after the 0, 0 is still an integer
	if len(s) == 0 { return false }
	for _, char := range s {
		if !strings.ContainsAny(strconv.QuoteRune(char), vOct) {
			return false
		}
	}
	return true
}
func checkValidDecimal(s string) bool {
	//define valid decimal characters
	vDec := "0123456789"
	//no leading characters to trim
	if len(s) == 0 { return false }
	for _, char := range s {
		if !strings.ContainsAny(strconv.QuoteRune(char), vDec) {
			return false
		}
	}
	return true
}