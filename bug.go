```go
package main

import "fmt"

func main() {
    var x int
    fmt.Println(x)
    x = append([]int{1,2,3}, 4)
}