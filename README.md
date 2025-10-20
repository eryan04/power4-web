<h1>🎮 Power4 Web (Puissance 4 en Go)</h1>

Mon projet web développé en Golang qui recrée le célèbre jeu Puissance 4 avec une interface HTML/CSS.

Le jeu se joue à deux joueurs, directement depuis le navigateur par defaut.

<h1>🧠 À propos du projet</h1>

L’application repose sur le package <mark>net/http</mark> pour gérer le serveur, les routes, ainsi que les templates HTML.

Le design minimaliste a été pensé pour être clair, rapide et aussi simple à jouer avec deux personnes sur un même écran.

<h1>🚀 Fonctionnalités principales</h1>

<h3>👥 Deux joueurs en local.</h3>

<h3>🧩 Trois niveaux de difficulté :</h3>

<pre>
  easy → Grille 6x7

  normal → Grille 6x9

  hard → Grille 7x8
</pre>

<h3>🏆 Détection automatique du vainqueur (ou match nul).</h3>

<h3>🔁 Bouton Rejouer avec les mêmes paramètres.</h3>

<h3>🆕 Bouton Nouvelle partie pour changer les noms et/ou la difficulté.</h3>

<h3>💅 Interface moderne sombre avec une touche de rouge.</h3>

<h1>🗂️ Structure du projet</h1>
<pre>
.
├── main.go                # Point d’entrée principal
├── go.mod                 # Dépendances Go
│
├── src/
│   ├── game.go            # Logique du jeu (grille, vérification, tours)
│   ├── handler.go         # Gestion des routes et templates
│   └── server.go          # Démarrage du serveur HTTP
│
├── pages/
│   ├── home.html          # Page d’accueil (formulaire)
│   └── play.html          # Page principale du jeu
│
├── templates/
│   ├── start.html         # Template inclus dans home.html
│   ├── tab.html           # Plateau de jeu dynamique
│   └── result.html        # Page de résultats
│
├── static/
│   ├── style.css          # Styles de la page d’accueil
│   └── play.css           # Styles du plateau et du résultat
</pre>

<h1>⚙️ Installation & Lancement</h1>

<h2>1️⃣ Cloner le dépôt :</h2>

<pre>
git clone https://github.com/eryan04/power4-web.git
cd power4-web
</pre>
<h2>2️⃣ Initialiser les modules Go :</h2>

<pre>
go mod tidy
</pre>
  
<h2>3️⃣ Lancer le serveur :</h2>

<pre>
go run main.go
</pre>
  
<h2>4️⃣ Jouer 🎲</h2>

Ouvre ton navigateur sur 👉 http://localhost:8080

<h1>🧩 Règles du jeu</h1>

Le but est d’aligner 4 pions consécutifs (horizontalement, verticalement ou en diagonale).

Les joueurs jouent à tour de rôle en cliquant sur les boutons "▼" pour choisir la colonne.

<h1>🧑‍💻 Auteur</h1>

Projet Power4 Web

Crée et developpé par Ryan Elqali. Etudiant à Ynov Lyon.
