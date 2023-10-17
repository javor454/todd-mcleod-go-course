# Array
- číslovaná sekvence elementů stejných typů
- nemění se její velikost
- používá se primárně pro GO internals => lepší použít **Slice**

# Slice
- podobný jako pole ale dynamická velikost
- má délku a kapacitu
- interně využívá pointer na array
  - při změně kapacity si runtime vytvoří nové pole o větší kapacitě a nakopíruje do něj prvky

# Map
- klíč => hodnota
- neseřazená skupina hodnot jednoho typu
- indexovaná unikátním klíčem jednoho typu
- rychlý lookup hodnot

# Struct
- KOMPOZITNÍ typ seskupuje hodnoty různých typů dohromady
- je možné je zanořovat
- inner type promotion - property z vnořené struktury se zpropagují do vnější struktury
- anonymní strukty
- nevytvářím třídy, ale TYPY
  - proč vytvářet typy?:
    - pokud jim chci přidat chování
    - pokud jimi chci dokumentovat kód
  - např. time Duration int64
- nevytvářím instance, ale HODNOTY
- pojmenované typy vs anonymní typy
  - anonymní jsou neurčité - kompilátor s nima pracuje flexibilně
  - do pojmenovaného typu můžu přiřadit kompatibilní anonymní typ
- pokud chci aby byl struct co nejrychlejší, je dobré seřadit typy dle velikosti (v bytech)
