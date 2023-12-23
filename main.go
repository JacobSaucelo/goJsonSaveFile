package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

type ItemType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SaveDataType struct {
	Items []ItemType `json:"Items"`
}

func main() {
	folderName := "data"

	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err := os.Mkdir(folderName, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
	}

	fPath := filepath.Join(folderName, "savefile.json")
	saveData, err := readSaveFile(fPath)
	if err != nil {
		fmt.Println("save file doesnt exists, creating a new one...")
		// fmt.Println("Error reading save file:", err)
		generateSaveFile()
		// return
	}

	addItem := ItemType{
		ID:   rand.Intn(100),
		Name: "blue gem karambit",
	}
	saveData.Items = append(saveData.Items, addItem)
	err = saveSaveFile(fPath, saveData)
	if err != nil {
		fmt.Println("error saving file")
		return
	}

}

func readSaveFile(fPath string) (SaveDataType, error) {
	saveFile, errSave := os.Open(fPath)
	if errSave != nil {
		fmt.Println("Opening file error")
		return SaveDataType{}, errSave
	}
	defer saveFile.Close()

	decoder := json.NewDecoder(saveFile)
	var save SaveDataType
	errSave = decoder.Decode(&save)
	if errSave != nil {
		fmt.Println("Decoding file error")
		return SaveDataType{}, errSave
	}

	return save, nil
}

func saveSaveFile(fPath string, data SaveDataType) error {
	saveFile, errSave := os.Create(fPath)
	if errSave != nil {
		fmt.Println("Opening file error")
		return errSave
	}
	defer saveFile.Close()

	encoder := json.NewEncoder(saveFile)
	errSave = encoder.Encode(data)
	if errSave != nil {
		fmt.Println("Opening file error")
		return errSave
	}

	return nil
}

func generateSaveFile() {
	generate := []byte(`{"Items":[]}`)
	fileErr := os.WriteFile("data/savefile.json", generate, 0644)
	if fileErr != nil {
		panic(fileErr)
	}
	fmt.Println("Successfully generated store.json")
}
