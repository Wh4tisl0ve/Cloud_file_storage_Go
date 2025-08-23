package main

import (
	"fmt"

	"github.com/Wh4tisl0ve/Cloud_file_storage_Go/internal/config"
)

func main() {
	config := config.MustLoad()

	fmt.Println(config)
}
