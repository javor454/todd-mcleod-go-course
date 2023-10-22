# Method sets
- metody A přijímající semantiku hodnot a metoda B přijímající semantiku pointerů obě podporují zavolání nad hodnotou i pointerem
- pokud ale signaturu metody A a B uzavřu do interfacu I a chci využít I jako parametr funkce, která provolává obě metody, argument volání musí být pointer
