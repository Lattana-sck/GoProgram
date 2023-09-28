package methods

import (
	"CC1/types"
	"encoding/csv"
	"fmt"
	"os"
)

func WriteCSV(username string, data []types.ResponsData) error {
	folderName := fmt.Sprintf("github/%s", username)
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err := os.Mkdir(folderName, os.ModePerm)
		if err != nil {
			return err
		}
	}

	csvFileName := fmt.Sprintf("%s/%s.csv", folderName, username)

	csvFile, err := os.Create(csvFileName)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	headers := []string{
		"Nom de l'utilisateur",
		"Nom du répo",
		"Url du répo",
	}

	if err := writer.Write(headers); err != nil {
		return err
	}

	for _, repo := range data {
		row := []string{
			username,
			repo.Name,
			repo.CloneURL,
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	fmt.Printf("Données écrites dans %s\n", csvFileName)
	return nil
}
