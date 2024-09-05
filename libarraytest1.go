package libarraytest1

import (
	"fmt"
	"math/rand"
	"time"
)

// ArrayGenerate retourne un tuple de nombres aléatoires de la taille spécifiée.
func ArrayGenerate(size int) ([]int, error) {
	if size <= 0 {
		return nil, fmt.Errorf("la taille doit être un nombre entier positif")
	}

	rand.Seed(time.Now().UnixNano())
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(10000) + 1
	}
	return result, nil
}
