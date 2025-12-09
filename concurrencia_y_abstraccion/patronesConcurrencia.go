package concurrenciayabstraccion

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var urls = []string {
	"https://rickandmortyapi.com/api/character/1",
	"https://rickandmortyapi.com/api/character/2",
	"https://rickandmortyapi.com/api/character/3",
	"https://rickandmortyapi.com/api/character/4",
	"https://rickandmortyapi.com/api/character/5",
	"https://rickandmortyapi.com/api/character/6",
	"https://rickandmortyapi.com/api/character/7",
	"https://rickandmortyapi.com/api/character/8",
	"https://rickandmortyapi.com/api/character/9",
	"https://rickandmortyapi.com/api/character/10",
}

func PatronConcurrencia() {
	example_sin_concurrencia()
	example_con_concurrencia()
}

func fetchURLSinConcurrencia(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}

	defer resp.Body.Close()
	duracion := time.Since(start)
	fmt.Printf("Fetched %s en %v - Status %s\n", url, duracion, resp.Status)
}

func example_sin_concurrencia() {
	fmt.Println("EJEMPLO SIN CONCURRENCIA")
	start := time.Now()

	for _, url := range urls {
		fetchURLSinConcurrencia(url)
	}

	duracionTotal := time.Since(start)
	fmt.Printf("Tiempo total de la secuencia: %v \n\n", duracionTotal)
}

func fetchURLConConcurrencia(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Marca la goroutine como completada

	start := time.Now()
	resp, err := http.Get(url)
	
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}

	defer resp.Body.Close()
	duracion := time.Since(start)
	fmt.Printf("Fetched %s en %v - Status %s\n", url, duracion, resp.Status)
}

func example_con_concurrencia() {
	fmt.Println("EJEMPLO CON CONCURRENCIA")
	start := time.Now()
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchURLConConcurrencia(url, &wg)
	}

	wg.Wait()
	duracionTotal := time.Since(start)
	fmt.Printf("Tiempo total concurrente: %v \n", duracionTotal)
}