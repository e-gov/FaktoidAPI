package rahvafakt

// Stack is a simple stack implementation
type Stack []string

// Push adds an element to stack
func (s *Stack) Push(v string) {
    *s = append(*s, v)
}

// Pop removes an element from the stack
func (s *Stack) Pop() string {
    if len(*s) == 0 {
        return ""
    }
    res := (*s)[len(*s) - 1]
    *s = (*s)[:len(*s) - 1]
    return res
}

// Content returns the content of the stack
func (s *Stack) Content() *[]string {
    var c []string
    var i int
    
    for _, l := range *s{
        c = append(c, l)
        i++
    }
    return &c
}


// CountDots counts the leading dots of a string
func CountDots(s string) int{
    var c int 
    
    for _, r := range s {
        if r == '.'{
            c++   
        } else{
            break
        }
    }    
    
    return c
}