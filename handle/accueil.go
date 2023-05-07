package handle

import (
	"fmt"
	"forum/forum"
	"html/template"
	"net/http"
)

// Accueil gere l'affichage sur la page d'acceuil
func Accueil(w http.ResponseWriter, r *http.Request) {
	// definision des page a executer
	pageconnecte := template.Must(template.ParseFiles("./templates/accueilco.html"))
	pagenonconnecte := template.Must(template.ParseFiles("./templates/accueil.html"))
	// recuperation de de la de la session utilisateur
	session, err := forum.Store.Get(r, "forum")
	// gestion de l'erreur
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Vérifier si l'utilisateur est connecté
	pseudo, ok := session.Values["pseudo"].(string)
	// entré les valeur dans la structure utilisateur
	utilisateurs := forum.Utilisateurs{
		Pseudo: pseudo,
	}
	// entré les valeurs dans la structure envoie
	envoie := forum.Envoie{
		User: utilisateurs,
	}
	// si l'utilisateur connecté
	if ok {
		// executer la page connecte
		pageconnecte.Execute(w, envoie)
	}
	// si l'utilisateur pas connecté
	if !ok {
		// executer la page non connecté
		err := pagenonconnecte.Execute(w, envoie)
		if err != nil {
			fmt.Println(err)
		}
	}
}
