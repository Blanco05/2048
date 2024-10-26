package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Nome del file
	fileName := "valore.txt"

	// Leggi il valore dal file
	value, err := leggiValoreDaFile(fileName)
	if err != nil {
		fmt.Println("Errore nella lettura del file:", err)
		return
	}

	// Crea scanner per input dell'utente
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Inserisci 'minchiata' per aggiungere 0.01, o 'giusto' per sottrarre 0.01:")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	// Aggiorna il valore in base all'input
	if input == "minchiata" {
		value += 0.01
	} else if input == "giusto" {
		value -= 0.01
	} else {
		fmt.Println("Input non valido. Inserisci 'minchiata' o 'giusto'.")
		return
	}

	// Scrivi il nuovo valore nel file
	err = scriviValoreNelFile(fileName, value)
	if err != nil {
		fmt.Println("Errore nella scrittura del file:", err)
		return
	}

	// Stampa il valore aggiornato
	fmt.Printf("Il nuovo valore è: %.2f\n", value)
}

// Funzione per leggere il valore dal file
func leggiValoreDaFile(fileName string) (float64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		// Se il file non esiste, inizia da 100.00
		if os.IsNotExist(err) {
			return 100.00, nil
		}
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		valoreStr := scanner.Text()
		valore, err := strconv.ParseFloat(valoreStr, 64)
		if err != nil {
			return 0, err
		}
		return valore, nil
	}

	return 100.00, nil // Default a 100.00 se il file è vuoto
}

// Funzione per scrivere il nuovo valore nel file
func scriviValoreNelFile(fileName string, value float64) error {
	file, err := os.Create(fileName) // Sovrascrive il file
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%.2f", value))
	if err != nil {
		return err
	}

	return nil
}
