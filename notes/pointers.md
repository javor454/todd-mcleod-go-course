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

TODO 21 - Pointers: 007 Pointers, values, the stack & the heap
