package main

import (
	"log"

	"kpp.dev/minireddit/internal/infra/memory"
	"kpp.dev/minireddit/internal/presenters"
)

func main() {
	log.Println("server initialized")

	memoryPostRepository := memory.NewPostRepositoryMemory()
	log.Println("loaded memory post repository")

	log.Println("starting http presenter")
	presenters.HTTP(memoryPostRepository)
}
