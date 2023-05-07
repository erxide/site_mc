package handle

import (
	"fmt"
	"forum/forum"
	"html/template"
	"net/http"
	"os/exec"
	"strconv"
)

func NouveauServeurMc(w http.ResponseWriter, r *http.Request) {
	pageconnecte := template.Must(template.ParseFiles("./templates/nouveauservermc.html"))
	session, err := forum.Store.Get(r, "forum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pseudo, ok := session.Values["pseudo"].(string)
	if !ok {
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		return
	}
	var serveurExiste int
	err = forum.Bd.QueryRow("SELECT COUNT(*) FROM McServeur WHERE proprio = ?", pseudo).Scan(&serveurExiste)
	if err != nil {
		fmt.Println(err)
	}
	if serveurExiste > 0 {
		// Si l'utilisateur a déjà un serveur, rediriger vers la page de gestion de serveur
		http.Redirect(w, r, "/votreserveur", http.StatusSeeOther)
		return
	}
	utilisateurs := forum.Utilisateurs{
		Pseudo: pseudo,
	}
	// entré les valeurs dans la structure envoie
	envoie := forum.Envoie{
		User: utilisateurs,
	}
	pageconnecte.Execute(w, envoie)
	port, err := forum.Port()
	proprio := pseudo
	_, err = forum.Bd.Exec("INSERT INTO McServeur (proprio, port) VALUES (?, ?)", proprio, port)
	if err != nil {
		fmt.Println(err)
	}
	str := strconv.Itoa(port)
	cmd := exec.Command("/bin/sh", "/srv/projetleo/minecraftserver/script/copynewserver.sh", proprio, str)
	err = cmd.Run()
	if err != nil {
        	fmt.Println(err)
	}
	fmt.Println("Nouveau serveur : ", pseudo)
	cmd = exec.Command("/bin/sh", "/srv/projetleo/minecraftserver/script/runserveur.sh", proprio)
        err = cmd.Run()
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println("Nouveau serveur : ", pseudo)
	url := "/votreserveur"
	fmt.Fprintf(w, "<script>window.location.href = '%s';</script>", url)
}
