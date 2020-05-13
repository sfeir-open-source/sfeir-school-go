# Le langage

- Né en 2009 chez Google (après les processeurs multi-coeurs) et OSS
- Binaire compilé autoporteur (début plugin depuis Go 1.8)
- Garbage collector (sub millisecond pour 17 Go de heap)
- Pointeurs 😱
- Goroutines
  - Assimilable à un thread
  - Mais ce n’est **PAS** un thread ⇒ **beaucoup plus léger**
  - Multithreadées sur un pool de thread
- Channels
  - **Do not communicate by sharing memory; share memory by communicating.**
  - Synchronisation
  - Multiplexage (**select**)

Notes:
OFU

CSP (1977)
Communicating sequential processes
