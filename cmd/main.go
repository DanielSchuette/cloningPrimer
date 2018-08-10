package main

import cloningprimer "github.com/DanielSchuette/cloningPrimer"

func main() {
	// find forward primer with EcoRI restriction site and 5 random nucleotides
	cloningprimer.FindForward("CAATGTGAGCTTAGCCTGATCCGTAATCGTAAGT", "GAATTC", 1, 10, 5)
}
