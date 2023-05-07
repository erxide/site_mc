package handle

import (
	"forum/forum"
	"net/http"
)

// Deconnexion gere la deconnexion de l'utilisateur
func Deconnexion(w http.ResponseWriter, r *http.Request) {
	// recuperation de de la de la session utilisateur
	session, err := forum.Store.Get(r, "forum")
	// gestion de l'erreur
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Supprime la session
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	// gestion de l'erreur
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// redirection de l'utilisateur vers l'acceuil
	http.Redirect(w, r, "/accueil", http.StatusSeeOther)
}
