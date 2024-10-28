package main
import (
	"fmt"
	"encoding/csv"
	"bytes"
	"io"
	"os"
)

func main(){
	data, err := readCSV("files.csv")
	if err!= nil {
		fmt.Println("Erreur lors la lecture du fichier", err)
		return 
	}

	reader, err := parseCSV(data)
	if err!= nil {
		fmt.Println("Erreur lors de la création du CSV reader")
		return
	}
	processCSV(reader)
}

func readCSV(filename string) ([]byte, error){
	f,err := os.Open(filename)
	if err!=nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err!= nil {
		return nil, err
	}
	return data, nil
}

func parseCSV(data []byte) (*csv.Reader, error){
	file := csv.NewReader(bytes.NewReader(data))
	return file,nil
}

func processCSV(reader *csv.Reader){
	incorrectAnswers,correctAnswers := 0,0
	for {
		record, err := reader.Read()
		if err == io.EOF{
			break
		} else if err!= nil {
			fmt.Println("Error reading CSV:", err)
		}
		var inputUser string
		fmt.Println("Veuillez répondre au question suiivante")
		fmt.Println(record[0])
		fmt.Scanln(&inputUser)
		if inputUser == record[1]{
			fmt.Println("Correct")
			correctAnswers +=1
		}else if inputUser != "0" || inputUser == ""{
			fmt.Println("Incorect")
			incorrectAnswers +=1
		}
	}
	fmt.Printf(
		"Vous avez trouvé %v \n Vous n'avez pas trouvé %v", correctAnswers, incorrectAnswers)
}