# Client-Serveur-golang
Applications client et serveur écrites en Go pour mon stage chez Trace Software en 2015/2016

## Déroulement du développement

Le but de ce mini projet est de me familiariser avec le langage [Go](https://golang.org/). L'objectif est de développer deux applications : un client et un serveur en rajoutant au fur et à mesure plusieurs fonctionnalités.

### [Etape 1](https://github.com/Mistermatt007/Client-Serveur-golang/commit/64b3ddefb509777e8b1ccc10eac038fc1e648bf4)
Créer un client qui envoie un message à une application serveur, le message est inscrit en dur dans le code source.
Quant au serveur il reçoit le message et l'affiche.

### [Etape 2](https://github.com/Mistermatt007/Client-Serveur-golang/commit/3f63f71afb1a6cd9cd0e3f0b5123494cc950017c)
Maintenant le client envoie un message que l'utilisateur aura saisi au serveur qui l'affiche puis ils s'envoient ce même message à l'infini.

### [Etape 3](https://github.com/Mistermatt007/Client-Serveur-Golang/commit/5c70b360d1ddfafb248658f3cdaa55ca268a45ec)
Le serveur doit maintenant accepter un certain nombre de clients défini par une variable d'environnement, le client se génère un identifiant qu'il envoie au serveur pour que le serveur les identifie et puisse leur envoyer un message lorsque le serveur est plein.

### [Etape 3 bis](https://github.com/Mistermatt007/Client-Serveur-Golang/commit/a41703ef2a830aa6d1b02eb6c8da10cbb00c8056)
Le serveur doit fonctionner de façon continue et doit savoir gérer les clients qui se déconnectent pour accueillir de nouveau clients. Le code doit être plus commenté, plus facile à lire.

Faire en sorte que le clients se débrouillent seuls, par exemple, ils ouvrent un connexion, essaient de dialoguer avec le serveur et ferment la connexion, ceci un certain nombre de fois afin de simuler de vrai clients et ainsi mieux tester le serveur.
