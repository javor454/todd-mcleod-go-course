# Konkurence
- zpusob jak psat kod pomoci konkurentních design patternu
- zpusob jak strukturovat kod tak aby mohl potencialne pouzit paralelismus
- umoznuje pridat kodu moznost jak spustit nekolik ukolu zaroven nezavisle
- v GO je konkurence umoznena pomoci **gorutin**, go program muze vytvorit nekolik gorutin ktere bezi nezavisle kazda zpracovavajici jiny ukol
- komunikace = vymena hodnot
- synchronizace = zajisteni ze N gorutin je ve znamem stavu
- komunikace a synchronizace gorutin funguje pomoci kanalu, ktere umoznuji gorutinam vymenu dat a koordinaci spusteni

# Paralelismus
- spousteni kodu paralelne na nekolika jadrech procesoru
- GO ma nativni podporu paralelismu

# Select 
- switch pro konkurentni kod
- sledovani nekolika kanalu na komunikaci
  - pokud ma default case a zadny kanal neni pripraveny, vykonam default
  - pokud nemam default case, cekam na prvni pripraveny kanal
  - pokud jsou pripravene oba kanaly, system vybere nahodne jeden

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

# Gorutina
- neni vlakno

# WaitGroup
- udrzuje si pocet gorutin, na ktery ma pockat
- s kazdou novou gorutinou pocet inkrementuji
- na zacatku kazde gorutiny pocet dekrementuji
- kdyz potrebuji tak muzu zablokovat chod a pockat nez se vsechny gorutiny dokonci

# Kanaly
- slouzi pro kumunikaci a synchronizaci mezi gorutinami

## Unbuffered kanal
- ma kapacitu 0
- zajisti sync mezi Sender a Receiver gorutinou
- pokud Sender posle hodnotu pred tim nez je Receiver ready -> zablokuje dokud neni Receiver ready
- pokud je Receiver ready ale Sender nic neposlal -> zablokuje dokud Sender neposle hodnotu

## Buffered kanal
- ma kapacitu > 0
- dostavam hodnoty v poradi co jsem je poslal FIFO
- dobre v pripade kdy:
  - ma Sender a Receiver ruzne rychlosti
  - chci decouplovat synchronizaci mezi nimi
- umozni gorutine poslat vice hodnot bez potreby okamzite synchronizace


