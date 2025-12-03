# GoLang-WalkThrough

## Lecture et calcul concurrent avec goroutines

Dans cet exercice, vous allez écrire un programme en Go qui lit des nombres depuis un fichier texte et effectue des calculs en parallèle à l’aide des goroutines et des channels.

---

### Objectif

* Lire une liste de nombres depuis un fichier numbers.txt.
* Pour chaque nombre, lancer une goroutine qui calcule son carré.
* Utiliser un channel pour récupérer les résultats et les afficher dans l’ordre où ils arrivent.


---

### Rappel pédagogique

Goroutine : une fonction exécutée de manière concurrente avec d’autres fonctions.


Channel : un mécanisme de communication sécurisé entre goroutines.


Pourquoi ? → Ce modèle permet de paralléliser des tâches simples et de comprendre la base de la concurrence en Go.


---

### Questions pour aller plus loin :

Que se passe-t-il si vous avez 10 000 nombres ?


Comment limiter le nombre de goroutines pour éviter la surcharge ? (1 go routine pour chaque nombre, consommation de ressources élevée -> utilisation des sync.WaitGroup et des workers)


Quelle différence entre ce modèle et un traitement séquentiel ?

