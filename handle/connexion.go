package handle

import (
	"fmt"
	"forum/forum"
	"html/template"
	"net/http"
)

// Connexion gere la connexion le l'utilisateur
func Connexion(w http.ResponseWriter, r *http.Request) {
	// page est le fichier html a executer
	page := template.Must(template.ParseFiles("./templates/connexion.html"))
	// recuperation de de la de la session utilisateur
	session, err := forum.Store.Get(r, "forum")
	// gestion de l'erreur
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Vérifier si l'utilisateur est connecté
	_, ok := session.Values["pseudo"].(string)
	if ok {
		// Rediriger l'utilisateur vers la page d'acceuil s'il est connecté
		http.Redirect(w, r, "/acceuil", http.StatusSeeOther)
		return
	}
	// methode post
	if r.Method == "POST" {
		// recuperation des valeurs
		pseudo := r.FormValue("pseudo")
		mdp := r.FormValue("mdp")
		// Vérifier les identifiants de l'utilisateur
		test, err := forum.Check(pseudo, mdp)
		// gestion de l'erreur
		if err != nil {
			fmt.Println(err)
			message := "error"
			Message := forum.ErreurMessage{
				Message: message,
			}
			page.Execute(w, Message)
		}
		// Authentification réussie : créer une session pour l'utilisateur
		if test {
			// ecrire dans le terminl que la connexion de tel utilisateur est reusite
			fmt.Println("connexion reussite de ", pseudo)
			// sauvegarde du pseudo de l'utilisateur dans la session
			session.Values["pseudo"] = pseudo
			session.Save(r, w)
			// redirection de l'utilisateur vers acceuil
			http.Redirect(w, r, "/accueil", http.StatusFound)
		} else {
			// Authentification échouée : afficher un message d'erreur
			message := "mauvaise combinaison"
			Message := forum.ErreurMessage{
				Message: message,
			}
			page.Execute(w, Message)
		}
	} else {
		// exection de la page avec un message
		message := "entrez identifiant"
		Message := forum.ErreurMessage{
			Message: message,
		}
		page.Execute(w, Message)

	}
}
