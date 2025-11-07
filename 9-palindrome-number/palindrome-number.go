import "strconv" 

func isPalindrome(x int) bool {
    // gonna turn it into str so we can reverse it
    str := strconv.Itoa(x)
    // reverse the str
    runes := []rune(str)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    // compare if its the same , EZPZ
    reversedStr := string(runes) 
    return str == reversedStr    
}