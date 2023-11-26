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
- funkce bežící konkurentné s dalšími gorutinami ve stejném adresovém prostoru
- kdyz je vytvořena, má svůj segment paměti v zásobníku, paměť v haldě je sdílena mezi všemi gorutinami
- jsou multiplexovány (převedeny na vstup) několika vláken operačního systému a v případě, že je jedno zablokování (např. kvůli blokaci čekáním na I/O) jiné pokračují v běhu
- klíčové slovo `go` spustí volání funkce nebo metody ve stejném adresovém prostoru
- program nečeká na dokončení gorutiny
  - pokud program skončí, všechny gorutiny se ukončí
  - pokud má gorutina návratovou hodnotu, při dokončení 
- neni vlakno

# WaitGroup
- udrzuje si pocet gorutin, na ktery ma pockat
- s kazdou novou gorutinou pocet inkrementuji
- na zacatku kazde gorutiny pocet dekrementuji
- kdyz potrebuji tak muzu zablokovat chod a pockat nez se vsechny gorutiny dokonci

# Kanaly
- slouzi pro kumunikaci a synchronizaci mezi gorutinami
- poslani a prijmuti pres kanal se musi stat ve stejnou chvili -> kanaly jsou blokujici

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

# Race condition
- v případě že dvě nebo více gorutin konkurentné přistoupí ke sdílené proměnné a alespoň jeden z přístupů je zápis
- vede k neočekávanému chování programu jako crash nebo korupce paměti
- možné odhalit pomocí `go run -race main.go`

# Synchronizační primitivy

## Mutex
- slouží k uzamčení kódu, kód mezi zamčením a odemčením může spustit pouze jedna gorutina

## RWMutex
- umoznuje uzamknout pro cteni a psani zaroven

