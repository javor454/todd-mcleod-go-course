# FMT
- Println
- Sprintf
- Printf
- Stringer interface - implementuje vse co ma metodu string ktera vypise obsah

# IO
- Writer TODO

# OS
- Create - vytvori file, vrati pointer

# bytes
- Buffer - buffer bytů
  - Write nebo WriteString - zápisová metody
  - Read nebo ReadString - čtecí metody
  - Bytes - vrátí obsah bufferu jako byte slice
  - Reset - resetne do původního stavu
- NewBuffer - vytvoří nový bytes.Buffer
- NewBufferString - vytvoří nový bytes.Buffer ze stringu
