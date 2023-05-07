package forum

func PseudoCheck(pseudo string) (bool, error) {
	// Requête SELECT pour récupérer les noms d'utilisateurs existants
	rows, err := Bd.Query("SELECT pseudo FROM Utilisateurs")
	if err != nil {
		return false, err
	}
	defer rows.Close()

	// Parcours des résultats pour vérifier si le nom d'utilisateur est déjà utilisé
	for rows.Next() {
		var PseudoExistant string
		if err := rows.Scan(&PseudoExistant); err != nil {
			return false, err
		}
		if PseudoExistant == pseudo {
			return true, nil

		}
	}

	// Si on est arrivé ici, c'est que le nom d'utilisateur n'est pas déjà pris
	return false, nil
}
