package main

import (
    "fmt"
    //"sync"
    "time"
)

func main() {
    //stringa := "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"
    //carattere := 'c'
    go writeTwoTimes("T ")
    writeTwoTimes("M ")
    //conteggio := conta(stringa, carattere)
    //fmt.Printf("Il carattere '%c' compare %d volte nella stringa\n", carattere, conteggio)
}

func writeTwoTimes(s string){
    for i:=0; i<2; i++{
        time.Sleep(100*time.Millisecond)
        fmt.Print(s)
    }
}

//func conta(stringa string, carattere rune) int