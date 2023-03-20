package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const cocktailDBAPIKey = "1"

type cocktailResponse struct {
	Drinks []struct {
		IDDrink                     string      `json:"idDrink"`
		StrDrink                    string      `json:"strDrink"`
		StrDrinkAlternate           interface{} `json:"strDrinkAlternate"`
		StrTags                     string      `json:"strTags"`
		StrVideo                    interface{} `json:"strVideo"`
		StrCategory                 string      `json:"strCategory"`
		StrIBA                      string      `json:"strIBA"`
		StrAlcoholic                string      `json:"strAlcoholic"`
		StrGlass                    string      `json:"strGlass"`
		StrInstructions             string      `json:"strInstructions"`
		StrInstructionsES           interface{} `json:"strInstructionsES"`
		StrInstructionsDE           string      `json:"strInstructionsDE"`
		StrInstructionsFR           interface{} `json:"strInstructionsFR"`
		StrInstructionsIT           string      `json:"strInstructionsIT"`
		StrInstructionsZHHANS       interface{} `json:"strInstructionsZH-HANS"`
		StrInstructionsZHHANT       interface{} `json:"strInstructionsZH-HANT"`
		StrDrinkThumb               string      `json:"strDrinkThumb"`
		StrIngredient1              string      `json:"strIngredient1"`
		StrIngredient2              string      `json:"strIngredient2"`
		StrIngredient3              string      `json:"strIngredient3"`
		StrIngredient4              string      `json:"strIngredient4"`
		StrIngredient5              interface{} `json:"strIngredient5"`
		StrIngredient6              interface{} `json:"strIngredient6"`
		StrIngredient7              interface{} `json:"strIngredient7"`
		StrIngredient8              interface{} `json:"strIngredient8"`
		StrIngredient9              interface{} `json:"strIngredient9"`
		StrIngredient10             interface{} `json:"strIngredient10"`
		StrIngredient11             interface{} `json:"strIngredient11"`
		StrIngredient12             interface{} `json:"strIngredient12"`
		StrIngredient13             interface{} `json:"strIngredient13"`
		StrIngredient14             interface{} `json:"strIngredient14"`
		StrIngredient15             interface{} `json:"strIngredient15"`
		StrMeasure1                 string      `json:"strMeasure1"`
		StrMeasure2                 string      `json:"strMeasure2"`
		StrMeasure3                 string      `json:"strMeasure3"`
		StrMeasure4                 interface{} `json:"strMeasure4"`
		StrMeasure5                 interface{} `json:"strMeasure5"`
		StrMeasure6                 interface{} `json:"strMeasure6"`
		StrMeasure7                 interface{} `json:"strMeasure7"`
		StrMeasure8                 interface{} `json:"strMeasure8"`
		StrMeasure9                 interface{} `json:"strMeasure9"`
		StrMeasure10                interface{} `json:"strMeasure10"`
		StrMeasure11                interface{} `json:"strMeasure11"`
		StrMeasure12                interface{} `json:"strMeasure12"`
		StrMeasure13                interface{} `json:"strMeasure13"`
		StrMeasure14                interface{} `json:"strMeasure14"`
		StrMeasure15                interface{} `json:"strMeasure15"`
		StrImageSource              string      `json:"strImageSource"`
		StrImageAttribution         string      `json:"strImageAttribution"`
		StrCreativeCommonsConfirmed string      `json:"strCreativeCommonsConfirmed"`
		DateModified                string      `json:"dateModified"`
	} `json:"drinks"`
}

func cocktailHandler(w http.ResponseWriter, r *http.Request) {
	// Envoyer une requête GET à l'API CocktailDB pour récupérer les cocktails commençant par la lettre 'a'
	resp, err := http.Get(fmt.Sprintf("https://www.thecocktaildb.com/api/json/v1/1/search.php?f=a&api_key=%s", cocktailDBAPIKey))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Lire le corps de la réponse en tant que JSON
	var cocktails cocktailResponse
	err = json.NewDecoder(resp.Body).Decode(&cocktails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Envoyer la réponse au format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cocktails)
}

func main() {
	http.HandleFunc("/cocktails", cocktailHandler)
	http.ListenAndServe(":8080", nil)
}