package main

import (
    "github.com/nathan-osman/go-rpigpio"
    "time"
    "flag"
)

func main() {
    //simple test to blink an LED n times on pin_num as specified by command line args
    pin_num := flag.Int("pin_num", 17, "an int")
    times := flag.Int("times", 10, "an int")
    flag.Parse()

    p, err := rpi.OpenPin(*pin_num, rpi.OUT)
    if err != nil {
        panic(err)
    }
    defer p.Close()

    for i := 0; i < *times; i++ {
        p.Write(rpi.HIGH)
        time.Sleep(300 * time.Millisecond)
        p.Write(rpi.LOW)
        time.Sleep(100 * time.Millisecond)
    }
}
