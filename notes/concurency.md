# Konkurence
- zpusob jak psat kod pomoci konkurentních design patternu
- umoznuje pridat kodu moznost jak spustit nekolik ukolu zaroven nezavisle
- v GO je konkurence umoznena pomoci **gorutin**, go program muze vytvorit nekolik gorutin ktere bezi nezavisle kazda zpracovavajici jiny ukol
- komunikace a synchronizace gorutin funguje pomoci kanalu, ktere umoznuji gorutinam vymenu dat a koordinaci spusteni

# Paralelismus
- spousteni kodu paralelne na nekolika jadrech procesoru

# Select 
- switch pro konkurentni kod 

# For cykly
- 3 druhy
  - for init, podminka, post
  - for podminka
  - for

# For range
- iterace nad rozsahem hodnot
- podpora pro slice (vycet stejnych hodnot) `[]int{1,2,3}`
- podpora pro mapu (mapa klicu a hodnot) `map[string]int{"Ahoj": 1, "Cus":  2}`
  - nevrati hodnoty v zadnem poradi
