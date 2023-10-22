# Bytes
- konverze stringu na slice bytů pomocí `[]byte("kamos")` zpet pomoci `string([]byte("kamos"))`

# Byte buffer
- region paměti sloužící k dočasnému uložení sekvence bytů
- poskytuje datovou strukturu pro efektivní manipulaci se sekvencemi bytů
- umožňuje číst a zapisovat do a z bufferu
- užitečné pro serializaci, síťovou komunikaci, file I/O, efektivní manipulaci se stringy

# interface{} nebo any
- hodnota jakéhokoliv typu

# Generika
- func SortStableFunc[S ~[]E, E any](x S, cmp func(a, b E) int)
  - E any - genericky parametr E muze byt jakykoliv typ
  - S ~[]E - genericky parametr S muze byt pole nebo slice (~) typu E
