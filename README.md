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
Crez un serveur de réponse en Go afin de répondre aux requêtes du client.

---

## Installation de Go

1. Téléchargez et installez Go depuis le site officiel : https://go.dev/dl/ (choisissez la dernière version stable pour votre système d'exploitation, par exemple **go1.25.5**)

2. Supprimez les anciennes versions de Go si nécessaire :

```bash
sudo rm -rf /usr/local/go
```

3. Extrayez l'archive téléchargée et installez Go (adaptez le nom du fichier si besoin) :

```bash
sudo tar -C /usr/local -xzf go1.25.5.linux-amd64.tar.gz
``` 

4. Ajoutez Go à votre variable d'environnement PATH :

- Si vous utilisez bash :

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

- Si vous utilisez zsh :

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
source ~/.zshrc
```

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

Ouvrez votre navigateur et rendez-vous sur : http://localhost:8080/


