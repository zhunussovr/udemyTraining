package main

import "fmt"

func main () {
   switch "Daniel" {
   case "Daniel":
       fmt.Println("Wassup Daniel")
   case "Mehdi":
       fmt.Println("Wassup Mehdi")
   case "Henny":
       fmt.Println("Wassup Henny")
   default:
       fmt.Println("Have you no friends?")

   }
}
