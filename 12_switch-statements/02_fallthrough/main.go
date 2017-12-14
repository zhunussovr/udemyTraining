package main

import "fmt"

func main () {
   switch "Henny" {
   case "Daniel":
       fmt.Println("Wassup Daniel")
       fallthrough
   case "Mehdi":
       fmt.Println("Wassup Mehdi")
   case "Henny":
       fmt.Println("Wassup Henny")
       fallthrough
   case "Marcus":
       fmt.Println("Wassup Marcus")
   }
}
