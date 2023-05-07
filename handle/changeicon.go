package handle

import (
	"forum/forum"
	"html/template"
	"log"
	"net/http"
)

// ChangeIcon est la fonction pour changer la couleur de l'icon
func ChangeIcon(w http.ResponseWriter, r *http.Request) {
	// page est le fichier html a executer
	page := template.Must(template.ParseFiles("./templates/changeicon.html"))
	// recuperation de de la de la session utilisateur
	session, err := forum.Store.Get(r, "forum")
	// gestion de l'erreur
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Vérifier si l'utilisateur est connecté
	pseudo, ok := session.Values["pseudo"].(string)
	if !ok {
		// Rediriger l'utilisateur vers la page de connexion s'il n'est pas connecté
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	// si requette post
	if r.Method == "POST" {
		// definision de IconColor
		IconCouleur := r.FormValue("icon")
		// Preparation de la requette sql pour changer la valeur de icon de l'utilisateur
		stmt, err := forum.Bd.Prepare("UPDATE Utilisateurs SET icon = ? WHERE pseudo = ?")
		// gestion de l'erreur
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		// exectution de la requette avec les valeur specifié
		_, err = stmt.Exec(IconCouleur, pseudo)
		// gestion de l'erreur
		if err != nil {
			log.Fatal(err)
		}
		// redirection de l'utilisateur vers l'accueil
		http.Redirect(w, r, "/accueil", http.StatusFound)
	}
	// execution de la page
	page.Execute(w, r)
}
