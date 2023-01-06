package os_study

import (
	"log"
	"os"
)

func StudyEnv() {
	username := os.Getenv("username")
	log.Printf("username %s \n", username)

	args := os.Args
	log.Printf("args %v", args)
}
