package handle

import (
	"fmt"
	"forum/forum"
	"html/template"
	"log"
	"net/http"
)

// ChangeMdp est la fonction pour changer le mot de passe
func ChangeMdp(w http.ResponseWriter, r *http.Request) {
	// page est le fichier html a executer
	page := template.Must(template.ParseFiles("./templates/changermdp.html"))
	// recuperation de de la de la session utilisateur
	session, err := forum.Store.Get(r, "forum")
	// gestion de l'erreur
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Vérifier si l'utilisateur est connecté
	pseudo, ok := session.Values["pseudo"].(string)
	// recuperation info utlisisateur
	_, _, _, _, _, icon, err := forum.ObtenirInfoUtilisateur(pseudo)
	Utilisateur := forum.Utilisateurs{
		Icon: icon,
	}

	if !ok {
		// Rediriger l'utilisateur vers la page de connexion s'il n'est pas connecté
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	// requette post
	if r.Method == "POST" {
		// recuperation des valeurs entre par l'utilisateur
		AncienMdp := r.FormValue("AncienMdp")
		NouveauMdp := r.FormValue("NouveauMdp")
		NouveauMdpCheck := r.FormValue("NouveauMdpCheck")
		// test pour savoir si l'utisateur a bien rentré son mot de passe
		test, err := forum.Check(pseudo, AncienMdp)
		// gestion de l'erreur
		if err != nil {
			fmt.Println(err)
		}
		// test ok bon ancien mot de passe
		if test {
			// si l'utilisateur a bien entré deux fois le meme mot de passe
			if NouveauMdp == NouveauMdpCheck {
				// hashage du nouveau mot de passe
				HashMdp, err := forum.HashMdp(NouveauMdp)
				// gestion de l'erreur
				if err != nil {
					log.Fatal(err)
				}
				// presparation de la requette sql pour changer le mot de passe de l'utilisateur
				stmt, err := forum.Bd.Prepare("UPDATE Utilisateurs SET mdp = ? WHERE pseudo = ?")
				// gestion de l'erreur
				if err != nil {
					log.Fatal(err)
				}
				defer stmt.Close()
				// execution de la requette sql
				_, err = stmt.Exec(HashMdp, pseudo)
				// gestion de l'erreur
				if err != nil {
					log.Fatal(err)
				}
				// redirection vers l'accueil
				http.Redirect(w, r, "/acceuil", http.StatusFound)
			} else {
				// si erreur alors executer la page avec le message d'erreur
				messageerreur := "Nouveau mot de passe et nouveau mot de passe ne sont pas les mêmes"
				message := forum.ErreurMessage{
					Message: messageerreur,
				}
				envoi := forum.Envoie{
					User:    Utilisateur,
					Message: message,
				}

				page.Execute(w, envoi)
			}
		} else {
			// si erreur alors executer la page avec le message d'erreur
			messageerreur := "Mauvais ancien mot de passe "
			message := forum.ErreurMessage{
				Message: messageerreur,
			}
			envoi := forum.Envoie{
				User:    Utilisateur,
				Message: message,
			}

			page.Execute(w, envoi)
		}
	} else {
		// sinon executer la page avec le message
		messageerreur := "Veuillez bien remplir tout les champs"
		message := forum.ErreurMessage{
			Message: messageerreur,
		}
		envoi := forum.Envoie{
			User:    Utilisateur,
			Message: message,
		}

		page.Execute(w, envoi)
	}

}
