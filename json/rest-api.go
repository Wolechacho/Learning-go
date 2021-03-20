package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

// Todo struct
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func CallApi() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(responseObject.Pokemon[i].Species.Name)
	}
}

func Post() {
	fmt.Println("2. Performing Http Post...")
	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	jsonReq, err := json.Marshal(todo)
	resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("%+v\n", todoStruct)
}
