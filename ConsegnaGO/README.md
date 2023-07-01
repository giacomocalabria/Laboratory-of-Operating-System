# Progetti GO - Sistemi Operativi UNIPD- Consegna

## Progetto 1 - Contatore di un carattere in una stringa

Il programma conta il numero di volte in cui un determinato carattere compare in una stringa, avviando una goroutine per ogni carattere della stringa. 

Il programma utilizza un meccanismo di sincronizzazione di tipo WaitGroup e un channel per evitare che più goroutine accedano contemporaneamente alla variabile condivisa che contiene il risultato.

## Progetto 2 - Agenzia di noleggio auto

Il programma simula il comportamento di un'agenzia di noleggio auto, che gestisce un insieme di auto di diversi tipi e le prenotazioni dei clienti. 

Il programma utilizza un channel per evitare che più goroutine accedano contemporaneamente alle variabili condivise che contano quante e quali auto sono state noleggiate. Utilizza inoltre un meccanismo di sincronizzazione di tipo WaitGroup per evitare che il programma termini prima che tutte le goroutine abbiano terminato.

## Progetto 3 - Produzione di torte

Il programma simula il comportamento di una pasticceria che produce torte e garantisce la sincronizzazione delle tre fasi di preparazione.

