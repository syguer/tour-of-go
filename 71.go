package main

import (
    "code.google.com/p/go-tour/tree"
    "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {   
    if t.Left != nil {
        Walk(t.Left, ch)
    }
    ch <- t.Value
    if t.Right != nil {
        Walk(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1:= make(chan int)
    go func(){
        Walk(t1, ch1)
        close(ch1)
    }()
    ch2:= make(chan int)
    go func(){
        Walk(t2, ch2)
        close(ch2)
    }()
    for a := range ch1 {
        b := <-ch2
        if a != b {
            return false
        }
    }
    return true
}

func main() {
    ch:= make(chan int)
    go func(){
        Walk(tree.New(1), ch)
        close(ch)
    }()
    for v := range ch {       
       fmt.Print(v)      
    }
    result := Same(tree.New(1), tree.New(1))
    fmt.Println(result)
    result = Same(tree.New(1), tree.New(2))
    fmt.Println(result)
}
