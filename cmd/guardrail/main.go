package main

import (
	"fmt"
	"guardrail/internal"
	"log"

	"guardrail/internal"
)

func main() {

	fmt.Println("🛡️ Guardrail Engine Starting...")

	// Path to the file Teraform created
	err := internal.CheckPolicy("terraform/policy_report.txt")
	if err != nil {
		log.Fatal(err)
	}

}
