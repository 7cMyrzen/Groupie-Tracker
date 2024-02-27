package mypackage

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//--------------------------------------------------------------------------------------------------\\
//------------------ Adresse IP de l'ordinateur pour l'affichage dans le terminal ------------------\\
//--------------------------------------------------------------------------------------------------\\

func GetIP() (string, error) {
	// Obtenez toutes les interfaces réseau de la machine
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("Erreur lors de la récupération des interfaces réseau: %v", err)
	}

	// Parcourez chaque interface pour obtenir les adresses IP
	for _, iface := range interfaces {
		// Vérifiez si l'interface est une carte Wi-Fi en vérifiant le nom
		if strings.Contains(iface.Name, "Wi-Fi") {
			addrs, err := iface.Addrs()
			if err != nil {
				return "", fmt.Errorf("Erreur lors de la récupération des adresses IP pour l'interface %s: %v", iface.Name, err)
			}

			// Parcourez chaque adresse IP associée à l'interface
			for _, addr := range addrs {
				// Vérifiez si l'adresse est de type IPv4
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					return ipNet.IP.String(), nil
				}
			}
		}
	}

	return "", fmt.Errorf("Aucune adresse IPv4 trouvée pour la carte Wi-Fi")
}

//----------------------------------------------------------------------------------------------------\\
//------------------ Modifier l'affichage de la date pour la lisibilite sur le site ------------------\\
//----------------------------------------------------------------------------------------------------\\

func TraduireDate(date time.Time) string {
	mois := map[int]string{
		1:  "janvier",
		2:  "février",
		3:  "mars",
		4:  "avril",
		5:  "mai",
		6:  "juin",
		7:  "juillet",
		8:  "août",
		9:  "septembre",
		10: "octobre",
		11: "novembre",
		12: "décembre",
	}
	return fmt.Sprintf("%d %s %d", date.Day(), mois[int(date.Month())], date.Year())
}
