package main

// lets map them :
var romanValues = map[byte]int{
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}
func romanToInt(s string) int {
    // Start with the value of the last symbol.
    total := romanValues[s[len(s)-1]]
    // Iterating backwards
    for i := len(s) - 2; i >= 0; i-- {
        currentVal := romanValues[s[i]]
        prevVal := romanValues[s[i+1]] 
        if currentVal < prevVal {
            total -= currentVal
        } else {
            total += currentVal
        }
    }
    return total
}
