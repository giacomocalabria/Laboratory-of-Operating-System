# Progetti GO - Sistemi Operativi - Consegna Finale

Giacomo Calabria - Matricola 2007964 - Sistemi Operativi A.A. 22/23

## Progetto 1 - Contatore di un carattere in una stringa

Il programma utilizza la concorrenza per contare il numero di volte in cui un determinato carattere compare in una stringa.

Il programma crea una goroutine per ogni caratteree della stringa. Questo permette di eseguire il conteggio dei caratteri in modo concorrente, sfruttando i vantaggi della parallelizzazione.

Utilizza un channel `conta` per tenere traccia del numero di occorrenze del carattere cercato, viene usato dalle varie goroutine per accedere in mutua esclusione alla variabile condivisa che contiene il risultato. La funzione `increaseNumber` viene utilizzata per incrementare il conteggio del carattere cercato nel channel.

Il programma utilizza un meccanismo di sincronizzazione di tipo WaitGroup per evitare che il programma termini prima che tutte le goroutine abbiano terminato. Il WaitGroup viene incrementato per ogni goroutine creata e decrementato quando la goroutine termina. Dopo il programma chiude il channel per evitare eventuali deadlock.

### Progetto 1.1 - Contatore di caratteri in una stringa

Questa versione del programma è una mia variazione del primo progetto, in cui vengono contate tutte le occorrenze dei caratteri presenti nella stringa, e non solo un carattere specifico.

Viene utilizzata una mappa di tipo `map[rune]int` per tenere traccia del numero di occorrenze di ogni carattere. La mappa contiene solo i caratteri che compaiono almeno una volta nella stringa. La mappa è contenuta all' interno di un channel `conta` per far si che le goroutine accedano in mutua esclusione alla mappa. La funzione `increaseNumber` è stata adattata per incrementare il conteggio del carattere cercato nella mappa. Se il carattere non è presente nella mappa, viene aggiunto con valore iniziale 0, quindi viene comunque incrementato. Il canale è stato inizializzato con una mappa vuota.

## Progetto 2 - Agenzia di noleggio auto

Il programma simula il comportamento di un'agenzia di noleggio auto, che gestisce un insieme di auto di diversi tipi e le prenotazioni dei clienti.

All' inizio il programma inizializza la lista dei clienti e un a lista di veicoli disponibili.

Viene utilizzato un channel `conta` per tenere traccia del numero di auto noleggiate per tipo e per garantire che le variabili condivise vengano accedute in mutua esclusione. La funzione `increaseNumber` viene utilizzata per incrementare il conteggio nel channel del tipo di auto noleggiata (parametro `key` della funzione). Il channel è di tipo mappa `map[int] int` e viene inizializzato con una mappa vuota.

La funzione `noleggia` viene utilizzata per simulare il noleggio di un'auto da parte di un cliente. Viene scelto casualmente un veicolo tra quelli disponobili e viene incrementato il contatore del tipo di veicolo scelto, utilizzando il canale. Viene stampato il nome del cliente e il tipo di veicolo noleggiato.

La funzione `stampa` alla fine del programma stampa il numero di auto noleggiate per tipo.

Infine, utilizza un meccanismo di sincronizzazione di tipo WaitGroup per evitare che il programma termini prima che tutte le goroutine abbiano terminato.

## Progetto 3 - Produzione di torte

Il programma simula la produzione di torte in una pasticceria, coinvolgendo tre pasticceri che lavorano in modo indipendente ma sincronizzando le varie fasi di preparazione.

Ogni pasticcere è rappresentato da una una funzione `pasticciere` che viene eseguita in una goroutine. Ogni funzione utilizza i canali per la sincronizzazione delle fasi. In particolare per la sincornizzazione tra i pasticceri vengono utilizzati i canali:

* `spazi1`: Per contare gli spazi liberi per il pasticcere #1.
* `spazi2`: Per contare gli spazi liberi per il pasticcere #2.
* `TorteCucinate`: Per tracciare le torte che sono state cucinate e sono pronte per la guarnizione.
* `TorteGuarnite`: Per tracciare le torte che sono state guarnite e sono pronte per la decorazione.
* `TorteFinite`: Per tracciare le torte che sono state completamente decorate e sono quindi finite.ù

Nel main vengono creati i canali, avviate le goroutine che rappresentano i tre pasticceri e viene utilizzato un ciclo `for` per aspettare che tutte e 5 le torte siano finite, osservando la dimensione del canale `TorteFinite`. Infine viene stampato il messaggio di fine produzione.

## Progetto 4 - Simulatore di trading

Il programma simula un ciclo di trading di valute in un mercato fittizio, generato casualmente.

La funzione `simulateMarketData`, eseguita in goroutine, si occupa di generare casualmente i dati di mercato per le tre coppie di valute. Ogni secondo vengono generate nuove quotazioni per ciascuna coppia e inviate sul canale `market`. Ogni valore varia all'interno di uno specifico range.

La funzione `selectPair`, eseguita in goroutine, è responsabile della selezione e dell'esecuzione delle operazioni tra le coppie di valute. In base a condizioni specificate. Utilizza una selezione `select` per gestire le condizioni di vendita in parallelo.

Nel main del programma, vengono creato il canale per comunicare i dati di mercato, viene creato il Mutex e vengono avviate le goroutine. La linea 108: `time.Sleep(time.Millisecond)` garantisce che il canale abbia dati di mercato prima di iniziare a eseguire le operazioni.

La fine del ciclo viene viene gestita verificando dapprima che il Mutex sia libero, e quindi che tutte le transazioni sianon completate. Infine viene stampato il messaggio di fine simulazione.
