# Client-Serveur-golang
Applications client et serveur écrites en Go pour mon stage chez Trace Software en 2015/2016

## Déroulement du développement

Le but de ce mini projet est de me familiariser avec le langage [Go](https://golang.org/). L'objectif est de développer deux applications : un client et un serveur en rajoutant au fur et à mesure plusieurs fonctionnalités.

### [Etape 1](https://github.com/Mistermatt007/Client-Serveur-golang/commit/64b3ddefb509777e8b1ccc10eac038fc1e648bf4)
Créer un client qui envoie un message à une application serveur, le message est inscrit en dur dans le code source.
Quant au serveur il reçoit le message et l'affiche.

### [Etape 2](https://github.com/Mistermatt007/Client-Serveur-golang/commit/3f63f71afb1a6cd9cd0e3f0b5123494cc950017c)
Maintenant le client envoie un message que l'utilisateur aura saisi au serveur qui l'affiche puis ils s'envoient ce même message à l'infini.

### Etape 3
Le serveur doit maintenant accepter un certain nombre de clients défini par une variable d'environnement, le serveur donne un numéro à chaque client et leur envoie. Le client ne fait qu'afficher ce numéro.
