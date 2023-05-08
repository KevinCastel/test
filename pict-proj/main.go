package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// URL de l'image à télécharger
	url := "https://images.unsplash.com/photo-1529118319834-474bdaa70346"

	// Téléchargement de l'image
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Création d'un fichier pour écrire l'image
	file, err := os.Create("image.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Copie du contenu de la réponse HTTP dans le fichier
	_, err = io.Copy(file, response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("L'image a été téléchargée et enregistrée avec succès.")
}
