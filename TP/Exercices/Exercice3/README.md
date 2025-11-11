# GoLang-WalkThrough

## Serveur de chat concurrent avec Echo

```html
../Assets/demo-exo3.mp4
```

Dans cet exercice, vous allez implémenter la partie serveur d’une application de chat en Go.
Le code HTML du client est déjà fourni pour la structure et l’interface, vous n’avez donc qu’à vous concentrer sur la logique serveur.

Objectif

* Créer un serveur HTTP avec Echo qui gère un chat en temps réel.
* Utiliser WebSockets pour la communication bidirectionnelle.
* Exploiter les goroutines et channels pour gérer la concurrence.
* Diffuser les messages à tous les clients connectés.

---

### Énoncé

1. Implémentez un serveur HTTP avec Echo.
2. Ajoutez une route /ws qui gère les connexions WebSocket.
3. Chaque client doit pouvoir :
    * Se connecter au serveur.
    * Envoyer des messages.
    * Recevoir les messages des autres clients.
4. Gérez la concurrence proprement :
    * Utilisez un channel pour la diffusion des messages.
    * Synchronisez avec un mutex pour la liste des clients.


Bonus :

* Ajouter des pseudos.
* Limiter le nombre de connexions.

---

### Commandes d'exécution

```bash
go run 3.go
```

Puis ouvrez le fichier client.html dans votre navigateur et connectez-vous au serveur (par défaut localhost:8080). N'hésitez dans la suite de l'exercice à ouvrir plusieurs pages afin d'avoir plusieurs clients connectés.