package forum

// LES DIFFERENTES STRUCTURES UTILISES DANS LE PROJET

type ErreurMessage struct {
	Message string
}

type Utilisateurs struct {
	ID     int
	Pseudo string
	Mdp    string
	Prenom string
	Nom    string
	Mail   string
	Age    int
	Icon   string
}

type McServeur struct {
	ID      int
	Proprio string
	Port    int
}

type Envoie struct {
	User      Utilisateurs
	Mcserveur McServeur
	Message   ErreurMessage
}
