package main

import (
    "fmt"
    "time"
)

func main() {
    go writeTwoTimes("T ")
    writeTwoTimes("M ")
}

func writeTwoTimes(s string){
    for i:=0; i<2; i++{
        time.Sleep(100*time.Millisecond)
        fmt.Println(s)
    }
}