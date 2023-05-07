package handle

import (
	"forum/forum"
	"net/http"
	"fmt"
	"os/exec"
)

func DlConfServeurVpn(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(" your user : " + pseudo)
	cmd := exec.Command("/bin/expect", "/srv/projetleo/minecraftserver/script/generate.exp", pseudo)
	//err = cmd.Run()
	//if err != nil {
        //        fmt.Println(err)
        //}
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Script execution failed with error: %v\nOutput: %s\n", err, output)
	} else {
		fmt.Printf("Script output: %s\n", output)
	}
	cmd = exec.Command("/bin/sh", "/srv/projetleo/minecraftserver/script/transferkey.sh", pseudo)
        output, err = cmd.CombinedOutput()
        if err != nil {
                fmt.Printf("Script execution failed with error: %v\nOutput:  %s\n", err, output)
        } else {
                fmt.Printf("Script output: %s\n", output)
        }
	cmd = exec.Command("/bin/sh", "/srv/projetleo/minecraftserver/script/make_config.sh", pseudo)
	output, err = cmd.CombinedOutput()
        if err != nil {
                fmt.Printf("Script execution failed with error: %v\nOutput:  %s\n", err, output)
        } else {
                fmt.Printf("Script output: %s\n", output)
        }
	referer := r.Header.Get("Referer")
	filePath := "/etc/openvpn/client/" + pseudo + "/" + pseudo + ".ovpn"
	fmt.Println("your filepath" + filePath)
	w.Header().Set("Content-Disposition", "attachment; filename="+pseudo+".ovpn")
	http.ServeFile(w, r, filePath)
	http.Redirect(w, r, referer, http.StatusFound)
}
	
