package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {

    f, err := os.Open("numbers.txt")

    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        var count int=0
        number, err := strconv.Atoi(scanner.Text())
        if err != nil {
            panic(err)
        }
        if number < 2 {
		    fmt.Println(number,"is not prime")

	    }
        for i:=2;i<number;i++{
            if number%i==0{
                count++
            }

        }
        if count==0{
            fmt.Println(number,"is prime")
        }
        if count>1{
            fmt.Println(number,"is not prime")
        }

    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}