package main

import "fmt"

func main () {
   
   myFriendName := "Henny"
   switch  {
   case len(myFriendName) == 2:
       fmt.Println("Wassup Daniel")
   case myFriendName == "Mehdi":
       fmt.Println("Wassup Mehdi")
   case myFriendName == "Henny":
       fmt.Println("Wassup - ", myFriendName)
   default:
       fmt.Println("Have you no friends?")

   }
}
