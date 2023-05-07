package main

import (
	"handle"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// initialisation du fichier assets pour pouvoir afficher le css et les images en front
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	// liste des routes http
	http.HandleFunc("/", handleSlash)
	http.HandleFunc("/accueil", handleAccueil)
	http.HandleFunc("/connexion", handleConnexion)
	http.HandleFunc("/deconnexion", handleDeconnexion)
	http.HandleFunc("/changeicon", handleChangeIcon)
	http.HandleFunc("/changemotdepasse", handleChangeMdp)
	http.HandleFunc("/test", handleTest)
	http.HandleFunc("/enregistrement", handleEnregistrement)
	http.HandleFunc("/nakar", handleNakar)
	http.HandleFunc("/votreserveur", handleVotreServeur)
	http.HandleFunc("/dlconfserveurvpn", handleDlConfServeurVpn)
	http.HandleFunc("/nouveauserv", handleNouveauServeurMc)
	// Écris dans le terminal, si le serveur a démarré, l'url du serveur avec le port
	log.Println("Serveur démarré sur http://localhost:8080")
	// Démarre le serveur sur le port 8080
	err := http.ListenAndServe(":8080", nil)
	// Si il y a une erreur
	if err != nil {
		// Stoppé le programme et écrire l'erreur dans le terminal
		log.Fatal(err)
	}
}

// Fonction handleSlash pour la route /
func handleSlash(w http.ResponseWriter, r *http.Request) {
	// redirection de l'utilisateur vers la route /accueil
	http.Redirect(w, r, "/accueil", http.StatusSeeOther)
}

// Fonction handleAccueil pour la route /accueil
func handleAccueil(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction Accueil dans le dossier forum
	handle.Accueil(w, r)
}

// Fonction handleConnexion pour la route /connexion
func handleConnexion(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction Connexion dans le dossier forum
	handle.Connexion(w, r)
}

// Fonction handleDeconnexion pour la route /connexion
func handleDeconnexion(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction deconnexion dans le dossier forum
	handle.Deconnexion(w, r)
}

// Fonction handleChangeIcon pour la route /connexion
func handleChangeIcon(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction changeicon dans le dossier forum
	handle.ChangeIcon(w, r)
}

// Fonction handleChangeMdp pour la route /connexion
func handleChangeMdp(w http.ResponseWriter, r *http.Request) {
	// appel de la fonction changemdp dans le dossier forum
	handle.ChangeMdp(w, r)
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	handle.Test(w, r)
}

func handleEnregistrement(w http.ResponseWriter, r *http.Request) {
	handle.Enregistrement(w, r)
}

func handleNakar(w http.ResponseWriter, r *http.Request) {
	handle.Nakar(w, r)
}

func handleVotreServeur(w http.ResponseWriter, r *http.Request) {
	handle.VotreServeur(w, r)
}

func handleDlConfServeurVpn(w http.ResponseWriter, r *http.Request) {
	handle.DlConfServeurVpn(w, r)
}

func handleNouveauServeurMc(w http.ResponseWriter, r *http.Request) {
	handle.NouveauServeurMc(w, r)
}
