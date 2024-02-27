// // handlers/groupes.go
package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type ArtistG struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"`
	Locations    string   `json:"locations"`
}

type ArtistHTML struct {
	ID            int    `json:"id"`
	Image         string `json:"image"`
	Name          string `json:"name"`
	MembersNumber int    `json:"members"`
	CreationDate  int    `json:"creationDate"`
	FirstAlbum    string `json:"firstAlbum"`
}

func GroupesHandler(w http.ResponseWriter, r *http.Request) {

	// URL de l'API des artistes
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"

	// Effectuer la requête HTTP GET
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Erreur lors de la requête HTTP:", err)
		return
	}
	defer response.Body.Close()

	// Lire le corps de la réponse
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du corps de la réponse:", err)
		return
	}

	// Convertir le corps de la réponse en une liste d'artistes
	var artists []ArtistG
	err = json.Unmarshal(body, &artists)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du JSON:", err)
		return
	}

	// Créer un modèle HTML pour les artistes
	artistsHTML := make([]ArtistHTML, len(artists))
	for i, artist := range artists {
		artistsHTML[i] = ArtistHTML{
			ID:            artist.ID,
			Image:         artist.Image,
			Name:          artist.Name,
			MembersNumber: len(artist.Members),
			CreationDate:  artist.CreationDate,
			FirstAlbum:    artist.FirstAlbum,
		}
	}

	// Créer un modèle HTML pour la page
	accueilFilePath := "templates/groupes.html"
	accueilFile, err := template.ParseFiles(accueilFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier HTML : %s", err), http.StatusInternalServerError)
		return
	}

	// Exécuter le modèle HTML
	err = accueilFile.Execute(w, artistsHTML)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML : %s", err), http.StatusInternalServerError)
		return
	}
}
