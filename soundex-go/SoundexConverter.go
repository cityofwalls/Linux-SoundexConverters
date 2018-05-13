package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

var conversions = map[string]int{
    "a": 0, "b": 1, "c": 2, "d": 3, "e": 0, "f": 1, "g": 2, "h": 0, "i": 0, "j": 2, "k": 2, "l": 4, "m": 5,
    "n": 5, "o": 0, "p": 1, "q": 2, "r": 6, "s": 2, "t": 3, "u": 0, "v": 1, "w": 0, "x": 2, "y": 0, "z": 2,
}

func main() {
    names := os.Args[1:]
    soundexNames := make([]string, len(names))
    
    // For each name in names: O(# of names)
    // 1. Replace all instances "w" and "h" with "" O(# of letters per name)
    
    // 2. Convert all letters other than first to their soundex number and remove duplicates O(# of letters per name)
    
    // 3. Replace all vowels with "" O(# of letters per name)
    
    // 4. Cut of name at 4 chars, or if len(name) < 4, pad with 0's, then replace current name in names with name[:4] O(1)
    
    for i, name := range names {
        noH := strings.Replace(name, "h", "", -1)
        noW := strings.Replace(noH, "w", "", -1)
        
        noDupes := removeDuplicateLetters(noW)
        
        noVowels := string(noDupes[:1])
        for _, letter := range noDupes[1:] {
            if letter != '0' {
                noVowels += string(letter)
            }
        }
        
        tooSmall := 4 - len(noVowels)
        var formatted string
        
        if tooSmall <= 0 {
            formatted = noVowels[:4]
        } else {
            formatted = noVowels
            for n := 0; n < tooSmall; n++ {
                formatted += "0"
            }
        }
        
        soundexNames = append(soundexNames[:i], formatted)
    }
    
    for _, name := range soundexNames {
        fmt.Printf("%s ", name)
    }
    fmt.Println()
}

func removeDuplicateLetters(name string) string {
    firstLetter := name[:1]
    convertedName := string(firstLetter)
    for _, letter := range name[1:] {
        convertedName += strconv.Itoa(conversions[string(letter)])
    }
    
    noDupes := string(firstLetter)
    previousLetter := strconv.Itoa(conversions[strings.ToLower(string(firstLetter))])
    for i, letter := range convertedName[1:] {
        if i > 0 {
            previousLetter = string(convertedName[i])
        }
        if !(string(letter) == string(previousLetter)) {
            noDupes += string(letter)
        }
    }
    return noDupes
}