package rahvafakt

type Stack []string

func (s *Stack) Push(v string) {
    *s = append(*s, v)
}

func (s *Stack) Pop() string {
    if len(*s) == 0 {
        return ""
    }
    res := (*s)[len(*s) - 1]
    *s = (*s)[:len(*s) - 1]
    return res
}

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