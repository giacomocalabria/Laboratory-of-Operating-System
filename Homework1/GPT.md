```go
package main

import (
  "fmt"
  "sync"
)

func main() {
  // Stringa di input
  input := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."

  // Carattere da cercare
  character := 'o'

  // Creazione di un WaitGroup per la sincronizzazione delle goroutine
  var wg sync.WaitGroup

  // Creazione di un channel per il conteggio totale dei caratteri corrispondenti
  count := make(chan int)

  // Iterazione sui caratteri della stringa
  for _, char := range input {
    // Avvio di una goroutine per ogni carattere
    wg.Add(1)
    go func(c rune) {
      defer wg.Done()

      // Verifica se il carattere corrisponde al carattere cercato
      if c == character {
        // Invio del carattere corrispondente al channel di conteggio
        count <- 1
      }
    }(char)
  }

  // Avvio di una goroutine per la chiusura del channel di conteggio
  go func() {
    wg.Wait()
    close(count)
  }()

  // Somma dei conteggi ricevuti dal channel
  totalCount := 0
  for c := range count {
    totalCount += c
  }

  // Stampa del conteggio totale
  fmt.Printf("Il carattere '%c' compare %d volte nella stringa.\n", character, totalCount)
}
```





c
