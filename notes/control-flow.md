# control flow
- pořadí ve kterém jsou spuštěny instrukce programu - příkazy
- druhy
  - sekvence
  - podmínky
  - cykly
- všechny go programy začínají v package main 
- v každém souboru se zdrojovým kódem může být init funkce na nastavení jakéhokoliv stavu je třeba

# Go runtime
- spravuje a plánuje exekuci go programů
- správa go rutin - lightweight vlákna která umožňují souběžně spustit programy - runtime vytváří a ničí gorutiny a plánuje exekuci na dostupných procesorech
- garbage kolekce - uvolňuje paměť kterou program již nepotřebuje
- správa paměti - spracuje alokaci a disalokaci paměti programu
- správa channelů - slouží pro komunikaci mezi gorutinami, runtime vytváří a ničí kanály a zaručuje doručení zpráv ve správném pořadi
- správa stacku - ukládá lokální proměnné a argumenty funkce, runtime spravuje alokaci a disalokaci místa na stacku pro každou gorutinu

# Stack
- sekce paměti sloužící pro ukládání proměnných lokálních pro funkci nebo gorutinu
  - pokaždé když je zavolána funkce nebo gorutina, vytvoří se nový rámec na stacku pro ukládání argumentů funkce, lokální proměnné a další data
- LIFO (poslední přidaná data jsou zpracována první)
- rychlejší než heap - alokaca a dealokace paměti je spravována automaticky GO runtimem

# Heap
- sekce paměti pro ukládání dlouhodobých proměnných do doby než jsou explicitně dealokovány nebo je program ukončen
- flexibilnější ale pomalejší - vyžaduje větší overhead pro alokaci a dealokaci paměti
- managuje garbage kolektor 

# Kompilátor
- přeloží lidsky čitelný kód na strojový (binární) kód - executable

# Podmínky
- == != < <= > >=

# Logické operátory
- ## || !

# Statement idiom (idiom = výraz)
- podmínce může předcházet krátký výraz, jehož scope je omezený do konce podmínky

