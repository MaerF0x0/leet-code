func isValid(s string) bool {
    stack := []rune{' '} // add an empty string to avoid invalid index panics
    
    for _, c := range []rune(s) { 
    
        switch c {
            case '{', '(',  '[':
            stack = append(stack, c)
            case '}', ')', ']':
            maybe := stack[len(stack)-1]
            if maybe == opposite(c) {
                stack = stack[:len(stack)-1]
                continue
            }
            return false
        }
    }
    return len(stack) == 1  // because of the " "       
}
    
func opposite(c rune) rune {
       switch c {
            case '}': return '{'
            case ')': return '('
            case ']': return '['
            default:
            return ' '
        }
}