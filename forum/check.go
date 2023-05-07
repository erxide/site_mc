package forum

import (
	"database/sql"
)

// check virifie les identifiant de l'uilisateur
func Check(pseudo string, mdp string) (bool, error) {
	// Récupérer le hash du mot de passe enregistré dans la base de données pour cet utilisateur
	row := Bd.QueryRow("SELECT mdp FROM Utilisateurs WHERE pseudo = ?", pseudo)
	var MdpHash string
	err := row.Scan(&MdpHash)
	if err != nil {
		if err == sql.ErrNoRows {
			// L'utilisateur n'existe pas dans la base de données
			return false, nil
		}
		return false, err
	}

	// Vérifier que le mot de passe fourni correspond au hash stocké dans la base de données
	ok, _ := CheckPasswordHash(mdp, MdpHash)
	if ok == true {
		// Le mot de passe fourni ne correspond pas au hash stocké dans la base de données
		return true, nil
	}

	// Les identifiants fournis sont valides
	return false, nil
}
