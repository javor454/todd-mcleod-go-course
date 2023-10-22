# FMT
- Println
- Sprintf
- Printf
- Stringer interface - implementuje vse co ma metodu string ktera vypise obsah

# IO
- Writer / Reader

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

# json
- Marshal - encode do jsonu
- Unmarshal - decode z jsonu
- Encode / Decode

# sort
- řazení
- Ints -> slices.Sort
- Strings -> slices.Sort

# slices
- metody nad slicy
- slices.Sort - serazeni slicu vzestupne (od nejmensiho po nejvetsi)
- slices.SortStableFunc - serazeni vzestupne pomoci clusury, ve ktere provolam compare

# x
- experimentalni pkg

# runtime
- GOOS - os info
- GOARCH - architektura CPU
- NumCPU - pocet jader
- NumGoroutine - pocet gorutin
