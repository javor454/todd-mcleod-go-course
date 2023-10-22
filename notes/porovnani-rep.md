- di - antipattern pouzivat, mel bysis slozit vsechny zavislosti sam
- recovery metoda je jedina co zachyti panic - chyba napr. pointer dereference
- zero logger
- nepouzivaji monorepa, pouze pokud to bude davat smysl kvuli domene
- validace struktury requestu - pres gin binduje request do moji preddefinovane struktury

- github.com/prometheus client golang
- michal posle odkaz na to pres co pousti coverage

- metoda `func New(...) *Market` -> vraci pointer na specificky typ s metodami (ne interface)
- interface definuju tam, kde se objekt pouziva pouze s metodama ktery potrebuje
- metody u typu maji pointer receiver 
