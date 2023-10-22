# Pointery (reference types)
- pointer je adresa v paměti
- umožňuje přímo manipulovat s pamětí
- operátor adresy & - získání adresy proměnné v paměti
- dereferencing operátor * - získání hodnoty uložené na adresse v paměti
- když předám proměnnou funkci, funkce dostane kopii proměnné - VŠECHNO V GO JE PASS BY VALUE - pokud funkce pozmění hodnotu, v originální proměnné se nic nezmění
  - pokud předám funkci pointer, funkce může pozměnit originální proměnnou  
  - pokud předám funkci referenční typ, změní originál
    - ref typy:
      - pointer
      - slice
      - mapa
      - channel
      - function
      - interface
- value semantika - predani dat jako hodnotu
  - jednodussi, bezpecnejsi, muze byt neefektivni v pripade velkych datovych struktur kvuli kopirovani dat
  - KDY POUZIT
    - default
    - pokud neni potreba modifikovat input
    - pokud jsou data malá
- pointer semantika - predani dat referenci
  - efektivnejsi a flexibilnejsi, vyzaduje opatrnou spravu sdileneho mutovatelneho stavu
  - KDY POUZIT
    - velká data 64 bytů a víc
    - pokud funkce nebo metoda potrebuje zmenit receivera nebo input
    - metody které updatují stav structu
    - KONZISTENCE - pokud má typ metodu s pointer semantikou, vsechny metody by meli mít metody s pointer semantikou

# Stack (zásobník) & Heap (halda)
- při použití pointer sémantiky je velká pravděpodobnost, že hodnoty escapnou na heap

## Value semantika a zásobník
- pokud fce přijme hodnotu (ne pointer) získá vlastní kopii té hodnoty
- tato kopie je typicky umístěna na stack, který je rychlý a nevyžaduje garbage collecting
- ve chvíli, kdy funkce vrátí, tak se paměť instantně uvolní

## Pointer semantika a halda
- pokud funkce získá nebo vrací pointer na lokální proměnnou, indikuje kompilátoru že tato hodnota se bude:
  - sdílet přes více gorutin
  - bude persistovat po návratu funkcí
- aby se data uchovala, kompilátor pro ně musí alokovat místo na haldě
- alokace paměti na haldě je náročnější a vyžaduje garbace collection

## Escape analýza
- určuje zda má být proměnná alokovaná v zásobníku nebo haldě
- zanalyzuje fci zda pointer na proměnnou se vrací nebo zda je proměnná uvnitř closure
  - pokud ano, proměnná půjde na haldu
  - pokud ne, proměnná půjde do zásobníku
- verbose escape analýza `go run -gcflags '-m' main.go`
  - vypíše informace
    ```
        ./main2.go:10:6: can inline functionA
        ./main2.go:11:13: inlining call to fmt.Println
        ./main2.go:7:7: inlining call to functionA
        ./main2.go:7:7: inlining call to fmt.Println
        ./main2.go:6:2: moved to heap: a
        ./main2.go:7:7: ... argument does not escape
        ./main2.go:10:12: leaking param: a
        ./main2.go:11:13: ... argument does not escape
    ```
