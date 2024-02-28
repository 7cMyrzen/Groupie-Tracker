// handlers/details.go
package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/mypackage"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"`
	Locations    string   `json:"locations"`
}

type ConcertsData struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Date struct {
	Date string `json:"date"`
}

type GroupDetails struct {
	ID           int           `json:"id"`
	Image        string        `json:"image"`
	Name         string        `json:"name"`
	Members      template.HTML `json:"members"`
	CreationDate template.HTML `json:"creationDate"`
	FirstAlbum   template.HTML `json:"firstAlbum"`
	Concerts     template.HTML `json:"concerts"`
	Locations    []string      `json:"locations"`
}

type LocationResponse struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesURL  string   `json:"dates"`
}

func getMembers(members []string) template.HTML {
	memberStr := ""
	for _, member := range members {
		memberStr = memberStr + "<br> - " + member
	}
	return template.HTML(memberStr)
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID du groupe depuis les paramètres de l'URL
	idStr := r.URL.Path[len("/groupe/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("ID de groupe invalide : %s", err), http.StatusBadRequest)
		return
	}

	// API URL pour les détails du groupe
	apiURL := "https://groupietrackers.herokuapp.com/api/artists/" + idStr

	// Effectuer la requête GET
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la requête GET : %s", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Lire le corps de la réponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la lecture du corps de la réponse : %s", err), http.StatusInternalServerError)
		return
	}

	//Declarer une variable de type Artist
	var artist Artist

	// Decoder le JSON
	err = json.Unmarshal(body, &artist)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la lecture du JSON : %s", err), http.StatusInternalServerError)
		return
	}

	members := getMembers(artist.Members)
	creationDate := template.HTML("<br> - le groupe " + artist.Name + " a été créé en " + strconv.Itoa(artist.CreationDate) + ".")
	FA, err := time.Parse("02-01-2006", artist.FirstAlbum)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la conversion de la date : %s", err), http.StatusInternalServerError)
		return
	}
	// Utiliser mypackage.TraduireDate pour obtenir la date formatée
	newFA := mypackage.TraduireDate(FA)
	firstAlbum := template.HTML("<br> - " + artist.Name + " a sorti son premier album le " + newFA + ".")

	relationAPI := artist.Relations
	//Effectuer la requête GET pour les relations du groupe
	resp, err = http.Get(relationAPI)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la requête GET pour les relations du groupe : %s", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Lire le corps de la réponse
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la lecture du corps de la réponse pour les relations du groupe : %s", err), http.StatusInternalServerError)
		return
	}

	//Declarer une variable de type ConcertsData
	var concertsData ConcertsData

	// Decoder le JSON des relations du groupe
	err = json.Unmarshal(body, &concertsData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la lecture du JSON pour les relations du groupe : %s", err), http.StatusInternalServerError)
		return
	}
	var concert string
	for location, dates := range concertsData.DatesLocations {
		for _, dateStr := range dates {
			// Utiliser time.Parse pour convertir la chaîne de date en objet time.Time
			date, err := time.Parse("02-01-2006", dateStr)
			if err != nil {
				http.Error(w, fmt.Sprintf("Erreur lors de la conversion de la date : %s", err), http.StatusInternalServerError)
				return
			}

			// Utiliser mypackage.TraduireDate pour obtenir la date formatée
			newDate := mypackage.TraduireDate(date)
			concert = concert + "<br> <span class=\"boldspan\">" + location + " :</span> " + newDate
		}
	}
	concerts := template.HTML(concert)

	locationurl := artist.Locations
	var locationlst []string
	//Effectuer la requête GET pour les locations du groupe
	resp, err = http.Get(locationurl)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la requête GET pour les locations du groupe : %s", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Lire le corps de la réponse
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la lecture du corps de la réponse pour les locations du groupe : %s", err), http.StatusInternalServerError)
		return
	}

	//Declarer une variable de type LocationResponse
	var locationResponse LocationResponse

	// Decoder le JSON des locations du groupe
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la lecture du JSON pour les locations du groupe : %s", err), http.StatusInternalServerError)
		return
	}
	locationlst = locationResponse.Locations

	// Créer un objet GroupDetails
	groupDetails := GroupDetails{
		ID:           id,
		Image:        artist.Image,
		Name:         artist.Name,
		Members:      members,
		CreationDate: creationDate,
		FirstAlbum:   firstAlbum,
		Concerts:     concerts,
		Locations:    locationlst,
	}

	// Créer un modèle pour les détails du groupe
	detailsFilePath := "templates/details.html"
	detailsFile, err := template.ParseFiles(detailsFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier HTML : %s", err), http.StatusInternalServerError)
		return
	}

	// Exécuter le modèle
	err = detailsFile.Execute(w, groupDetails)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML : %s", err), http.StatusInternalServerError)
		return
	}
}
