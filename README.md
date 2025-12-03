# GoLang-WalkThrough

## 🎯 Objectif du TP
Ce TP a pour but de vous familiariser avec Go, la programmation concurrente et la création d’applications réseau. Ce dernier aura la forme d'un CTF (Catpure The Flag), donc aura comme finalité de réussir à récupérer le flag, portant la forme FLAG{X}. Tout le TP se déroulera sur le serveur web du container docker. Il vous est demandé de programmer uniquement en Go, donc l'utilisation d'outils ou d'autres langages de programmation sera interdite.

Vous réaliserez 3 étapes successives :

---

### ✅ Challenge 1 : Fouilles en Base64

Lire le fichier volumineux contenant des chemins encodés en Base64.
Décoder, requêter et filtrer ces chemins pour trouver ceux qui mènent à la suite du challenge.

---

### ✅ Challenge 2 : Déceler les secrets

Casser des secrets dont le hash bcrypt est fourni.
Comparer une version séquentielle et une version parallélisée avec goroutines et workers.

--- 

### ✅ Challenge 3 : Le coffre

Accéder au coffre sécurisé.
Faites la rencontre d'un mystérieux client.
Répondez lui jusqu'à ce qu'il vous donne une récompense.
Crez un serveur de réponse à l'aide du framework Echo de Go.


---

## Exécution du projet avec Docker

### Construire l'image Docker

```bash
docker build -t ctf-go .
```

### Exécuter le container

```bash
docker run -p 8080:80 ctf-go
```

### Accéder au projet

Ouvrez votre navigateur et rendez-vous sur :

http://localhost:8080/

