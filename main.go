package main

import (
	"fmt"
	"groupie-tracker/handlers"
	"groupie-tracker/mypackage"
	"net/http"
	"text/template"
)

func main() {

	// Définissez les couleurs pour les messages de la console
	greenColor := "\033[32m"
	blueColor := "\033[34m"
	redColor := "\033[31m"
	defaultColor := "\033[0m"

	// Gestionnaire pour le répertoire /static/
	staticDir := "/static/"
	http.Handle(staticDir, http.StripPrefix(staticDir, http.FileServer(http.Dir("static"))))

	// Gestionnaire pour la route /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		accueilFilePath := "templates/accueil.html"
		accueilFile, err := template.ParseFiles(accueilFilePath)

		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier HTML : %s", err), http.StatusInternalServerError)
			return
		}

		err = accueilFile.Execute(w, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML : %s", err), http.StatusInternalServerError)
			return
		}
	})

	// Gestionnaire pour la route /groupes
	http.HandleFunc("/groupes", handlers.GroupesHandler)
	// Gestionnaire pour la route /groupe avec les details
	http.HandleFunc("/groupe/", handlers.DetailsHandler)

	// Récupérez l'adresse IP de la carte Wi-Fi pour afficher les adresses du site
	wifiIP, err := mypackage.GetIP()
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	// Affichez les adresses du site
	fmt.Println(greenColor, "Le site est disponible aux adresses suivantes :")
	fmt.Println()
	fmt.Println(blueColor, "Accueil :")
	fmt.Println(blueColor, "     http://localhost:8080")
	fmt.Println(blueColor, "     http://"+wifiIP+":8080")
	fmt.Println(blueColor, "Groupes :")
	fmt.Println(blueColor, "     http://localhost:8080/groupes")
	fmt.Println(blueColor, "     http://"+wifiIP+":8080/groupes", defaultColor)
	fmt.Println()
	fmt.Println(redColor, "Appuyez sur Ctrl+C pour arrêter le serveur", defaultColor)
	http.ListenAndServe(":8080", nil)
}
