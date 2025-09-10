

``` go
a := [3]int{0,1,2}
for i,a_copy := range a {
    a[i] = 0
    fmt.Printf(a_copy)
}

```