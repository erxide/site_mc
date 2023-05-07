package handle

import (
	"forum/forum"
	"net/http"
)

// SupPost gere la supression des post
func SupPost(w http.ResponseWriter, r *http.Request) {
	// extraction des valeurs
	err := r.ParseForm()
	// gestion de l'erreur
	if err != nil {
		http.Error(w, "Erreur lors de l'extraction des données de la requête", http.StatusInternalServerError)
		return
	}
	postID := r.FormValue("post_id")
	// requette sql pour supprimer le postes en question
	_, err = forum.Bd.Exec("DELETE FROM Postes WHERE id = ?", postID)
	// gestion de l'erreur
	if err != nil {
		http.Error(w, "Erreur lors de la mise à jour du nombre de likes", http.StatusInternalServerError)
		return
	}

	// Rediriger l'utilisateur vers la page précédente après avoir liké le poste
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}
