package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

    array := [50]int{r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21), r1.Intn(21)}

    fmt.Println(array)
   fmt.Println()
   fmt.Println()
   fmt.Println()
    Even :=0
    Odd := 0

  
    for i :=0; i < 50; i++ {
      if array[i]%2 == 0 {
      
     
      Even ++
     } else if array[i]%2 == 1 {
      
     
     Odd ++
    }
  } 
  fmt.Println("There are", Even, "even numbers")
  fmt.Println("There are", Odd, "odd numbers")
   
   fmt.Println()
   fmt.Println()

   one :=0
   for t :=0; t <50; t++{
    if array[t] ==1 {
      one ++
    }
   } 
   if one == 0 {
    fmt.Println("1: ")
   }
   if one == 1 {
    fmt.Println("1: 1")  
   }
   if one == 2 {
    fmt.Println("1: 1,1")
   }
   if one == 3 {
    fmt.Println("1: 1,1,1")
   }
   if one == 4 {
    fmt.Println("1: 1,1,1,1")
   }
   if one == 5 {
    fmt.Println("1: 1,1,1,1,1")
   }
   if one == 6 {
    fmt.Println("1: 1,1,1,1,1,1")
   }
   if one == 7 {
    fmt.Println("1: 1,1,1,1,1,1,1")
   }
   if one == 8 {
    fmt.Println("1: 1,1,1,1,1,1,1,1")
   }
   if one == 9 {
    fmt.Println("1: 1,1,1,1,1,1,1,1,1")
   }
   if one == 10 {
    fmt.Println("1: 1,1,1,1,1,1,1,1,1,1")
   }

   two :=0
   for t :=0; t <50; t++{
    if array[t] ==2 {
      two ++
    }
   } 
   if two == 0 {
    fmt.Println("2: ")
   }
   if two == 1 {
    fmt.Println("2: 2")  
   }
   if two == 2 {
    fmt.Println("2: 2,2")
   }
   if two == 3 {
    fmt.Println("2: 2,2,2")
   }
   if two == 4 {
    fmt.Println("2: 2,2,2,2")
   }
   if two == 5 {
    fmt.Println("2: 2,2,2,2,2")
   }
   if two == 6 {
    fmt.Println("2: 2,2,2,2,2,2")
   }
   if two == 7 {
    fmt.Println("2: 2,2,2,2,2,2,2")
   }
   if two == 8 {
    fmt.Println("2: 2,2,2,2,2,2,2,2")
   }
   if two == 9 {
    fmt.Println("2: 2,2,2,2,2,2,2,2,2")
   }
   if two == 10 {
    fmt.Println("2: 2,2,2,2,2,2,2,2,2,2")
   }

   three :=0
   for t :=0; t <50; t++{
    if array[t] ==3 {
      three ++
    }
   } 
   if three == 0 {
    fmt.Println("3: ")
   }
   if three == 1 {
    fmt.Println("3: 3")  
   }
   if three == 2 {
    fmt.Println("3: 3,3")
   }
   if three == 3 {
    fmt.Println("3: 3,3,3")
   }
   if three == 4 {
    fmt.Println("3: 3,3,3,3")
   }
   if three == 5 {
    fmt.Println("3: 3,3,3,3,3")
   }
   if three == 6 {
    fmt.Println("3: 3,3,3,3,3,3")
   }
   if three == 7 {
    fmt.Println("3: 3,3,3,3,3,3,3")
   }
   if three == 8 {
    fmt.Println("3: 3,3,3,3,3,3,3,3")
   }
   if three == 9 {
    fmt.Println("3: 3,3,3,3,3,3,3,3,3")
   }
   if three == 10 {
    fmt.Println("3: 3,3,3,3,3,3,3,3,3,3")
   }

   four :=0
   for t :=0; t <50; t++{
    if array[t] ==4 {
      four ++
    }
   } 
   if four == 0 {
    fmt.Println("4: ")
   }
   if four == 1 {
    fmt.Println("4: 4")  
   }
   if four == 2 {
    fmt.Println("4: 4,4")
   }
   if four == 3 {
    fmt.Println("4: 4,4,4")
   }
   if four == 4 {
    fmt.Println("4: 4,4,4,4")
   }
   if four == 5 {
    fmt.Println("4: 4,4,4,4,4")
   }
   if four == 6 {
    fmt.Println("4: 4,4,4,4,4,4")
   }
   if four == 7 {
    fmt.Println("4: 4,4,4,4,4,4,4")
   }
   if four == 8 {
    fmt.Println("4: 4,4,4,4,4,4,4,4")
   }
   if four == 9 {
    fmt.Println("4: 4,4,4,4,4,4,4,4,4")
   }
   if four == 10 {
    fmt.Println("4: 4,4,4,4,4,4,4,4,4,4")
   }
  
   five :=0
   for t :=0; t <50; t++{
    if array[t] ==5 {
      five ++
    }
   } 
   if five == 0 {
    fmt.Println("5: ")
   }
   if five == 1 {
    fmt.Println("5: 5")  
   }
   if five == 2 {
    fmt.Println("5: 5,5")
   }
   if five == 3 {
    fmt.Println("5: 5,5,5")
   }
   if five == 4 {
   fmt.Println("5: 5,5,5,5")
   }
   if five == 5 {
    fmt.Println("5: 5,5,5,5,5")
   }
   if five == 6 {
    fmt.Println("5: 5,5,5,5,5,5")
   }
   if five == 7 {
    fmt.Println("5: 5,5,5,5,5,5,5")
   }
   if five == 8 {
    fmt.Println("5: 5,5,5,5,5,5,5,5")
   }
   if five == 9 {
    fmt.Println("5: 5,5,5,5,5,5,5,5,5")
   }
   if five == 10 {
    fmt.Println("5: 5,5,5,5,5,5,5,5,5,5")
   }
  
   six :=0
   for t :=0; t <50; t++{
    if array[t] ==6 {
      six ++
    }
   } 
   if six == 0 {
    fmt.Println("6: ")
   }
   if six == 1 {
    fmt.Println("6: 6")  
   }
   if six == 2 {
    fmt.Println("6: 6,6")
   }
   if six == 3 {
    fmt.Println("6: 6,6,6")
   }
   if six == 4 {
    fmt.Println("6: 6,6,6,6")
   }
   if six == 5 {
    fmt.Println("6: 6,6,6,6,6")
   }
   if six == 6 {
    fmt.Println("6: 6,6,6,6,6,6")
   }
   if six == 7 {
    fmt.Println("6: 6,6,6,6,6,6,6")
   }
   if six == 8 {
    fmt.Println("6: 6,6,6,6,6,6,6,6")
   }
   if six == 9 {
    fmt.Println("6: 6,6,6,6,6,6,6,6,6")
   }
   if six == 10 {
    fmt.Println("6: 6,6,6,6,6,6,6,6,6,6")
   }
  
   seven :=0
   for t :=0; t <50; t++{
    if array[t] ==7 {
      seven ++
    }
   } 
   if seven == 0 {
    fmt.Println("7: ")
   }
   if seven == 1 {
    fmt.Println("7: 7")  
   }
   if seven == 2 {
    fmt.Println("7: 7,7")
   }
   if seven == 3 {
    fmt.Println("7: 7,7,7")
   }
   if seven == 4 {
    fmt.Println("7: 7,7,7,7")
   }
   if seven == 5 {
    fmt.Println("7: 7,7,7,7,7")
   }
   if seven == 6 {
    fmt.Println("7: 7,7,7,7,7,7")
   }
   if seven == 7 {
    fmt.Println("7: 7,7,7,7,7,7,7")
   }
   if seven == 8 {
    fmt.Println("7: 7,7,7,7,7,7,7,7")
   }
   if seven == 9 {
    fmt.Println("7: 7,7,7,7,7,7,7,7,7")
   }
   if seven == 10 {
    fmt.Println("7: 7,7,7,7,7,7,7,7,7,7")
   }
  
   eight :=0
   for t :=0; t <50; t++{
    if array[t] ==8 {
      eight ++
    }
   } 
   if eight == 0 {
    fmt.Println("8: ")
   }
   if eight == 1 {
    fmt.Println("8: 8")  
   }
   if eight == 2 {
    fmt.Println("8: 8,8")
   }
   if eight == 3 {
    fmt.Println("8: 8,8,8")
   }
   if eight == 4 {
    fmt.Println("8: 8,8,8,8")
   }
   if eight == 5 {
    fmt.Println("8: 8,8,8,8,8")
   }
   if eight == 6 {
    fmt.Println("8: 8,8,8,8,8,8")
   }
   if eight == 7 {
    fmt.Println("8: 8,8,8,8,8,8,8")
   }
   if eight == 8 {
    fmt.Println("8: 8,8,8,8,8,8,8,8")
   }
   if eight == 9 {
    fmt.Println("8: 8,8,8,8,8,8,8,8,8")
   }
   if eight == 10 {
    fmt.Println("8: 8,8,8,8,8,8,8,8,8,8")
   }
  
   nine :=0
   for t :=0; t <50; t++{
    if array[t] ==9 {
      nine ++
    }
   } 
   if nine == 0 {
    fmt.Println("9: ")
   }
   if nine == 1 {
    fmt.Println("9: 9")  
   }
   if nine == 2 {
    fmt.Println("9: 9,9")
   }
   if nine == 3 {
    fmt.Println("9: 9,9,9")
   }
   if nine == 4 {
    fmt.Println("9: 9,9,9,9")
   }
   if nine == 5 {
    fmt.Println("9: 9,9,9,9,9")
   }
   if nine == 6 {
    fmt.Println("9: 9,9,9,9,9,9")
   }
   if nine == 7 {
    fmt.Println("9: 9,9,9,9,9,9,9")
   }
   if nine == 8 {
    fmt.Println("9: 9,9,9,9,9,9,9,9")
   }
   if nine == 9 {
    fmt.Println("9: 9,9,9,9,9,9,9,9,9")
   }
   if nine == 10 {
    fmt.Println("9: 9,9,9,9,9,9,9,9,9,9")
   }
  
   ten :=0
   for t :=0; t <50; t++{
    if array[t] ==10 {
      ten ++
    }
   } 
   if ten == 0 {
    fmt.Println("10: ")
   }
   if ten == 1 {
    fmt.Println("10: 10")  
   }
   if ten == 2 {
    fmt.Println("10: 10,10")
   }
   if ten == 3 {
    fmt.Println("10: 10,10,10")
   }
   if ten == 4 {
    fmt.Println("10: 10,10,10,10")
   }
   if ten == 5 {
    fmt.Println("10: 10,10,10,10,10")
   }
   if ten == 6 {
    fmt.Println("10: 10,10,10,10,10,10")
   }
   if ten == 7 {
    fmt.Println("10: 10,10,10,10,10,10,10")
   }
   if ten == 8 {
    fmt.Println("10: 10,10,10,10,10,10,10,10")
   }
   if ten == 9 {
    fmt.Println("10: 10,10,10,10,10,10,10,10,10")
   }
   if ten == 10 {
    fmt.Println("10: 10,10,10,10,10,10,10,10,10,10")
   }
    
   eleven :=0
   for t :=0; t <50; t++{
    if array[t] ==11 {
      eleven ++
    }
   } 
   if eleven == 0 {
    fmt.Println("11: ")
   }
   if eleven == 1 {
    fmt.Println("11: 11")  
   }
   if eleven == 2 {
    fmt.Println("11: 11,11")
   }
   if eleven == 3 {
    fmt.Println("11: 11,11,11")
   }
   if eleven == 4 {
    fmt.Println("11: 11,11,11,11")
   }
   if eleven == 5 {
    fmt.Println("11: 11,11,11,11,11")
   }
   if eleven == 6 {
    fmt.Println("11: 11,11,11,11,11,11")
   }
   if eleven == 7 {
    fmt.Println("11: 11,11,11,11,11,11,11")
   }
   if eleven == 8 {
    fmt.Println("11: 11,11,11,11,11,11,11,11")
   }
   if eleven == 9 {
    fmt.Println("11: 11,11,11,11,11,11,11,11,11")
   }
   if eleven == 10 {
    fmt.Println("11: 11,11,11,11,11,11,11,11,11,11")
   }
    
   twelve :=0
   for t :=0; t <50; t++{
    if array[t] ==12 {
      twelve ++
    }
   } 
   if twelve == 0 {
    fmt.Println("12: ")
   }
   if twelve == 1 {
    fmt.Println("12: 12")  
   }
   if twelve == 2 {
    fmt.Println("12: 12,12")
   }
   if twelve == 3 {
    fmt.Println("12: 12,12,12")
   }
   if twelve == 4 {
    fmt.Println("12: 12,12,12,12")
   }
   if twelve == 5 {
    fmt.Println("12: 12,12,12,12,12")
   }
   if twelve == 6 {
    fmt.Println("12: 12,12,12,12,12,12")
   }
   if twelve == 7 {
    fmt.Println("12: 12,12,12,12,12,12,12")
   }
   if twelve == 8 {
    fmt.Println("12: 12,12,12,12,12,12,12,12")
   }
   if twelve == 9 {
    fmt.Println("12: 12,12,12,12,12,12,12,12,12")
   }
   if twelve == 10 {
    fmt.Println("12: 12,12,12,12,12,12,12,12,12,12")
   }
    
   thirteen :=0
   for t :=0; t <50; t++{
    if array[t] ==13 {
      thirteen ++
    }
   } 
   if thirteen == 0 {
    fmt.Println("13: ")
   }
   if thirteen == 1 {
    fmt.Println("13: 13")  
   }
   if thirteen == 2 {
    fmt.Println("13: 13,13")
   }
   if thirteen == 3 {
    fmt.Println("13: 13,13,13")
   }
   if thirteen == 4 {
    fmt.Println("13: 13,13,13,13")
   }
   if thirteen == 5 {
    fmt.Println("13: 13,13,13,13,13")
   }
   if thirteen == 6 {
    fmt.Println("13: 13,13,13,13,13,13")
   }
   if thirteen == 7 {
    fmt.Println("13: 13,13,13,13,13,13,13")
   }
  if thirteen == 8 {
    fmt.Println("13: 13,13,13,13,13,13,13,13")
   }
   if thirteen == 9 {
    fmt.Println("13: 13,13,13,13,13,13,13,13,13")
   }
   if thirteen == 10 {
    fmt.Println("13: 13,13,13,13,13,13,13,13,13,13")
   }
    
   fourteen :=0
   for t :=0; t <50; t++{
    if array[t] ==14 {
      fourteen ++
    }
   } 
   if fourteen == 0 {
    fmt.Println("14: ")
   }
   if fourteen == 1 {
    fmt.Println("14: 14")  
   }
   if fourteen == 2 {
    fmt.Println("14: 14,14")
   }
   if fourteen == 3 {
    fmt.Println("14: 14,14,14")
   }
   if fourteen == 4 {
    fmt.Println("14: 14,14,14,14")
   }
   if fourteen == 5 {
    fmt.Println("14: 14,14,14,14,14")
   }
   if fourteen == 6 {
    fmt.Println("14: 14,14,14,14,14,14")
   }
   if fourteen == 7 {
    fmt.Println("14: 14,14,14,14,14,14,14")
   }
   if fourteen == 8 {
    fmt.Println("14: 14,14,14,14,14,14,14,14")
   }
   if fourteen == 9 {
    fmt.Println("14: 14,14,14,14,14,14,14,14,14")
   }
   if fourteen == 10 {
    fmt.Println("14: 14,14,14,14,14,14,14,14,14,14")
   }
    
   fifteen :=0
   for t :=0; t <50; t++{
    if array[t] ==15 {
      fifteen ++
    }
   } 
   if fifteen == 0 {
    fmt.Println("15: ")
   }
   if fifteen == 1 {
    fmt.Println("15: 15")  
   }
   if fifteen == 2 {
    fmt.Println("15: 15,15")
   }
   if fifteen == 3 {
    fmt.Println("15: 15,15,15")
   }
   if fifteen == 4 {
    fmt.Println("15: 15,15,15,15")
   }
   if fifteen == 5 {
    fmt.Println("15: 15,15,15,15,15")
   }
   if fifteen == 6 {
    fmt.Println("15: 15,15,15,15,15,15")
   }
   if fifteen == 7 {
    fmt.Println("15: 15,15,15,15,15,15,15")
   }
   if fifteen == 8 {
    fmt.Println("15: 15,15,15,15,15,15,15,15")
   }
   if fifteen == 9 {
    fmt.Println("15: 15,15,15,15,15,15,15,15,15")
   }
   if fifteen == 10 {
    fmt.Println("15: 15,15,15,15,15,15,15,15,15,15")
   }
    
   sixteen :=0
   for t :=0; t <50; t++{
    if array[t] ==16 {
      sixteen ++
    }
   } 
   if sixteen == 0 {
    fmt.Println("16: ")
   }
   if sixteen == 1 {
    fmt.Println("16: 16")  
   }
   if sixteen == 2 {
    fmt.Println("16: 16,16")
   }
   if sixteen == 3 {
    fmt.Println("16: 16,16,16")
   }
   if sixteen == 4 {
    fmt.Println("16: 16,16,16,16")
   }
   if sixteen == 5 {
    fmt.Println("16: 16,16,16,16,16")
   }
   if sixteen == 6 {
    fmt.Println("16: 16,16,16,16,16,16")
   }
   if sixteen == 7 {
    fmt.Println("16: 16,16,16,16,16,16,16")
   }
   if sixteen == 8 {
    fmt.Println("16: 16,16,16,16,16,16,16,16")
   }
   if sixteen == 9 {
    fmt.Println("16: 16,16,16,16,16,16,16,16,16")
   }
   if sixteen == 10 {
    fmt.Println("16: 16,16,16,16,16,16,16,16,16,16")
  }
    
   seventeen :=0
   for t :=0; t <50; t++{
    if array[t] ==17 {
      seventeen ++
    }
   } 
   if seventeen == 0 {
    fmt.Println("17: ")
   }
   if seventeen == 1 {
    fmt.Println("17: 17")  
   }
   if seventeen == 2 {
    fmt.Println("17: 17,17")
   }
   if seventeen == 3 {
    fmt.Println("17: 17,17,17")
   }
   if seventeen == 4 {
    fmt.Println("17: 17,17,17,17")
   }
   if seventeen == 5 {
    fmt.Println("17: 17,17,17,17,17")
   }
   if seventeen == 6 {
    fmt.Println("17: 17,17,17,17,17,17")
   }
   if seventeen == 7 {
    fmt.Println("17: 17,17,17,17,17,17,17")
   }
   if seventeen == 8 {
    fmt.Println("17: 17,17,17,17,17,17,17,17")
   }
   if seventeen == 9 {
    fmt.Println("17: 17,17,17,17,17,17,17,17,17")
   }
   if seventeen == 10 {
    fmt.Println("17: 17,17,17,17,17,17,17,17,17,17")
   }
   
   eightteen :=0
   for t :=0; t <50; t++{
    if array[t] ==18 {
      eightteen ++
    }
   } 
   if eightteen == 0 {
    fmt.Println("18: ")
   }
   if eightteen == 1 {
    fmt.Println("18: 18")  
   }
   if eightteen == 2 {
    fmt.Println("18: 18,18")
   }
   if eightteen == 3 {
    fmt.Println("18: 18,18,18")
   }
   if eightteen == 4 {
    fmt.Println("18: 18,18,18,18")
   }
   if eightteen == 5 {
    fmt.Println("18: 18,18,18,18,18")
   }
   if eightteen == 6 {
    fmt.Println("18: 18,18,18,18,18,18")
   }
   if eightteen == 7 {
    fmt.Println("18: 18,18,18,18,18,18,18")
   }
   if eightteen == 8 {
    fmt.Println("18: 18,18,18,18,18,18,18,18")
   }
   if eightteen == 9 {
    fmt.Println("18: 18,18,18,18,18,18,18,18,18")
   }
   if eightteen == 10 {
    fmt.Println("18: 18,18,18,18,18,18,18,18,18,18")
   }
      
   nineteen :=0
   for t :=0; t <50; t++{
    if array[t] ==19 {
      nineteen ++
    }
   } 
   if nineteen == 0 {
    fmt.Println("19: ")
   }
   if nineteen == 1 {
    fmt.Println("19: 19")  
   }
   if nineteen == 2 {
    fmt.Println("19: 19,19")
   }
   if nineteen == 3 {
    fmt.Println("19: 19,19,19")
   }
   if nineteen == 4 {
    fmt.Println("19: 19,19,19,19")
   }
   if nineteen == 5 {
    fmt.Println("19: 19,19,19,19,19")
   }
   if nineteen == 6 {
    fmt.Println("19: 19,19,19,19,19,19")
   }
   if nineteen == 7 {
    fmt.Println("19: 19,19,19,19,19,19,19")
   }
   if nineteen == 8 {
    fmt.Println("19: 19,19,19,19,19,19,19,19")
   }
   if nineteen == 9 {
    fmt.Println("19: 19,19,19,19,19,19,19,19,19")
   }
   if nineteen == 10 {
    fmt.Println("19: 19,19,19,19,19,19,19,19,19,19")
   }
       
   twenty :=0
   for t :=0; t <50; t++{
    if array[t] ==20 {
      twenty ++
    }
   } 
   if twenty == 0 {
    fmt.Println("20: ")
   }
   if twenty == 1 {
    fmt.Println("20: 20")  
   }
   if twenty == 2 {
    fmt.Println("20: 20,20")
   }
   if twenty == 3 {
    fmt.Println("20: 20,20,20")
   }
   if twenty == 4 {
    fmt.Println("20: 20,20,20,20")
   }
   if twenty == 5 {
    fmt.Println("20: 20,20,20,20,20")
   }
   if twenty == 6 {
    fmt.Println("19: 20,20,20,20,20,20")
   }
   if twenty == 7 {
    fmt.Println("20: 20,20,20,20,20,20,20")
   }
   if twenty == 8 {
    fmt.Println("20: 20,20,20,20,20,20,20,20")
   }
   if twenty == 9 {
    fmt.Println("20: 20,20,20,20,20,20,20,20,20")
   }
   if twenty == 10 {
    fmt.Println("20: 20,20,20,20,20,20,20,20,20,20")
   }
   
}