package main

import (
    "fmt"
    "strconv"
)

var folder []file
type file []string

func main() {

    //just initializing some stuff here
    folder = make([]file, 10)
    for i := range folder {
      folder[i] = make([]string, 100)
      for j := range folder[i] {
	      folder[i][j] = "Some string" + strconv.Itoa(j)
      }
    }
    fileChan := make(chan file)
    lineChan := make(chan string)



    // you can adjust the limits in these 2 loops to control how many go routines are running

    for i := 0; i < 2 ; i++ {
      go processFiles(fileChan, lineChan)
    }
    for i := 0; i < 100; i++ {
      go processLines(lineChan)
    }

    for _, f := range folder {
	    fileChan <- f
    }
    // you'll want to keep it from exiting after this until its done
}

func processFiles(fileChan chan file, lineChan chan string) {
  for {
    select{
      case file := <-fileChan:
      for _, line := range file {
        lineChan <- line
      }
    }
  }
}

func processLines(lineChan chan string) {
  for {
    select{
      case line := <-lineChan:
      fmt.Println(line)
    }
  }
}