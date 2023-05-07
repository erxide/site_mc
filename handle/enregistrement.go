package handle

import (
	"fmt"
	"forum/forum"
	"html/template"
	"net/http"
)

// Enregistrement gere l'enregistrement des comptes dans la base de donnée
func Enregistrement(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("./templates/enregistrement.html"))
	session, err := forum.Store.Get(r, "forum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, ok := session.Values["pseudo"].(string)
	if ok {
		http.Redirect(w, r, "/accueil", http.StatusSeeOther)
		return
	}
	if r.Method == "POST" {
		pseudo := r.FormValue("pseudo")
		mdp := r.FormValue("mdp")
		mdp_conf := r.FormValue("mdp_conf")
		taken, err := forum.PseudoCheck(pseudo)
		if err != nil {
			fmt.Println(err)
			return
		}
		if taken {
			messageerror := "Utilisateur deja utilisé !"
			Message := forum.ErreurMessage{
				Message: messageerror,
			}
			page.Execute(w, Message)
			return
		}
		if mdp != mdp_conf {
			messageerror := "mot de passe pas concordant !"
			Message := forum.ErreurMessage{
				Message: messageerror,
			}
			page.Execute(w, Message)
			return
		}
		hashmdp, _ := forum.HashMdp(mdp)
		_, err = forum.Bd.Exec("INSERT INTO Utilisateurs (pseudo, mdp) VALUES (?, ?)", pseudo, hashmdp)
		// gestion de l'erreur
		if err != nil {
			fmt.Println(err)
		}
		// ecrire dans le terminal quand un compte est creer
		fmt.Println("Nouveau compte : ", pseudo)
		// Rediriger vers la page d'accueil'
		http.Redirect(w, r, "/accueil", http.StatusSeeOther)
	} else {
		// si il y a erreur alors executer la page avec le message d'erreur
		messageerror := "Entrez bien toute les informations"
		Message := forum.ErreurMessage{
			Message: messageerror,
		}
		page.Execute(w, Message)
	}
}
