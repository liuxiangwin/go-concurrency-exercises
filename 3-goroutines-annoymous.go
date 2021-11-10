package main
import "time"

func main() {
  go func (msg string){
	println(msg)
  }("Hello World!")

  time.Sleep(time.Second)
}

