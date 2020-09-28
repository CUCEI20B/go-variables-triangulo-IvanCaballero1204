package main

import "fmt"

func main()  {
  var base uint64;
  var altura uint64;

  fmt.Scan(&base);
  fmt.Scan(&altura)

  result := base * altura / 2
  fmt.Println(result)
}