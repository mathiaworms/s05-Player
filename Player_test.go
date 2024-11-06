package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPlayer(t *testing.T) {
	p4 := Player{
		Name:     "Patrick",
		Username: "THE_STAR",
	}
	input := "Patrick\nTHE_STAR\n"
	r, w, err := os.Pipe() // Créer un pipe pour rediriger l'entrée
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	// Simuler l'entrée avec un buffer
	go func() {
		w.Write([]byte(input)) // Écrire l'entrée simulée dans le pipe
		w.Close()              // Fermer l'écriture une fois terminé
	}()

	// Sauvegarder l'entrée standard originale
	originalStdin := os.Stdin
	os.Stdin = r // Rediriger os.Stdin vers le pipe

	// Appeler AddPlayer pour voir s'il crée le bon joueur
	ptest := AddPlayer()

	// Restaurer l'entrée standard originale après le test
	os.Stdin = originalStdin
	if ptest.Name != p4.Name || ptest.Username != p4.Username {
		t.Fatalf("Player loaded does not match expected values: got %+v, expected %+v", ptest, p4)
	}
	ptest.save()
	NameofFile := ptest.Name
	NameofFileWithExt := fmt.Sprintf("%s.yml", NameofFile)

	if _, err := os.Stat(NameofFileWithExt); err != nil {
		t.Fatalf("File doesn't exist exists\n got:nothing expected values: %d ", &NameofFileWithExt)
	}
	ptest.del()
	if _, err := os.Stat(NameofFileWithExt); err == nil {
		t.Fatalf("File doesn't exist exists\n got: %d expected values: nothing ", &NameofFileWithExt)
	}
}
