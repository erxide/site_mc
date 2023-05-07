package forum

import (
	"crypto/sha256"
	"encoding/hex"
)

// CheckPasswordHash vérifie si le mot de passe en clair correspond au hash stocké dans la base de données
func CheckPasswordHash(mdp string, hash string) (bool, error) {
	// Initialise la variable test par le hash de mdp
	test, err := HashMdp(mdp)
	// si HashMdp retourne une erreur
	if err != nil {
		// retourne faux
		return false, err
	}
	// si test est égal à hash
	if test == hash {
		// retourne vrai
		return true, nil
	}
	// sinon retourne dans tous les faux cas
	return false, err

}

// HashMdp hash le mot de passe en utilisant bcrypt
func HashMdp(mdp string) (string, error) {
	// Convertir le mot de passe en une série de bytes
	MdpBytes := []byte(mdp)

	// Calculer le hash SHA-256
	hashBytes := sha256.Sum256(MdpBytes)

	// Convertir le hash en une chaîne hexadécimale
	hashString := hex.EncodeToString(hashBytes[:])

	return hashString, nil
}
