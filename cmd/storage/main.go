package main

import (
	"fmt"

	"github.com/4klb/storage_ozon/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("hello world", st)
}
