# Funkce
- typ se stejnými možnostmi jako proměnné a typy
  - je možné je přiřadit k proměnné, předat jako argument funkci, vrátit jako návratovou hodnotu funkce
- význam:
  - abstrakce kódu
  - znovupoužitelnost
  - čitelnost

- formát (signature):
  - `func (receiver) identifier(parameters) (returns) { code }`
  - definujeme funkci s parametry
  - voláme funkci s argumenty - PASS BY VALUE

# Variadic parametr
- dynamické množšví parametrů
- variadic = can vary - může se lišit
- musí být poslední mezi parametry

# Defer
- odlož provolání funkce dokud funkce, v jejímž scopu se volání nachází neskončí (jakkoliv - return, došlo na konec, panic..)
- LIFO

# Metody
- pomocí receiveru přidám typu chování formou funkce - metody

# Interface
- definuje signaturu metod
- v GO, hodnoty můžou mít více typů
  - pokud má stejné metody, je stejný typ
- implicitni implementace typem

# Polymorfismus
- umožňuje hodnotám různých typů být považován za instanci společného typu
- pokud dva typy implementují stejnou metodu

# Wrapper funkce
- funk která poskytuje přídavnou úroveň abstrakce nebo funkcionality jiné existující funkci
- vhodné na:
  - logování
  - profilování
  - autentikace a autorizace
  - error handling

# Anonymní funkce
- funkce bez názvu kterou rovnou provolám
- signature anon funkce `func (p parameter(s)) (r return(s)) { code }`

# Func expression
- funkce kterou uložím do proměnné

# Callback
- předám funkci jako argument funkci
- signature `func doOperation(a, b int, f func(int, int) int) int { code }`

# Closure
- funkce, co v sobě uzavírá další funkci

# Recursion
- funkce volá sama sebe, dokud nenarazí na base case

# Dokumentace
- napíšu komentář nad implementaci funkce  
