package main 

import (
 "bufio"
 "fmt"
 "os"
 "strings"
)

func main(){
 scanner := bufio.NewScanner(os.Stdin)
 for scanner.Scan(){
  fmt.Println(strings.TrimSpace(scanner.Text()))
 }
}
