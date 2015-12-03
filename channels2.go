package main

func addData() {

    data := make(chan int)

    go func() {
        // get data from internet
        // ...
        data <- downloadNumber()
    }

    go func() {
        // get data from user
        // ...
        data <- inputNumber()
    }

    x := <-data
    y := <-data

    return x + y
}
