package piscine

func IsAlpha(str string) bool {
    r := []rune(str)
    for i := range r {
        if (runes[i-1] >= 33 && runes[i-1] <= 47) || (runes[i-1] >= 58 && runes[i-1] <= 64) || (runes[i-1] >= 91 && runes[i-1] <= 96) || (runes[i-1] >= 123 && runes[i-1] <= 126) || runes[i-1] == ' ' {
            return false
        }
    }
    return true
}
