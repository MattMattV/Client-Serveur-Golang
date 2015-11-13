# Client-Serveur-golang
Applications client et serveur écrites en Go pour mon stage chez Trace Software en 2015/2016

## Déroulement du développement

Le but de ce mini projet est de me familiariser avec le langage [Go](https://golang.org/). L'objectif est de développer deux applications : un client et un serveur en rajoutant au fur et à mesure plusieurs fonctionnalités.

### [Etape 1](https://github.com/Mistermatt007/Client-Serveur-golang/commit/1caaf88272593b718a067d4b63b660729f9611cc)
Créer un client qui envoie un message à une application serveur, le message est inscrit en dur dans le code source.
Quant au serveur il reçoit le message et l'affiche.

### [Etape 2](https://github.com/Mistermatt007/Client-Serveur-golang/commit/79c40449a6ab5d84ca2e49887f658ba72ba808ae)
Maintenant le client envoie le message que l'utilisateur à saisi au serveur qui l'affiche puis ils s'envoient ce même message à l'infini.

### Etape 3
Le serveur doit maintenant accepter un certain nombre de clients défini par une variable d'environnement, le serveur donne un numéro à chaque client et leur envoie. Le client ne fait qu'afficher ce numéro.
