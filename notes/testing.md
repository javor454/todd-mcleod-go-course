# Testy
- testy ziji ve stejnem packagi jako kod, ktery testuji
- jmena konvence `filename_test.go` kde `filename` je nazev souboru s testovanym kodem

# Testovaci funkce
- musi zacinat slovem `Test` a nasledovat slovem co zacina velkym pismenem
  - e.g. soubor `add_test.go` testuje soubor `add.go` a obsahuje funkci `func TestAdd(t *testing.T)` která testuje funkci `func Add(a,b int) int`
