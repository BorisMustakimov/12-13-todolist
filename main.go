package main

import (
	"fmt"
	"os"

	"github.com/BorisMustakimov/12-13-todolist/server"
)

func main() {
	a, err := server.New()
	if err != nil {
		fmt.Println("ошибка запуска приложения", err)
		os.Exit(1)
	}

	if err = a.Run(); err != nil {
		fmt.Println("ошибка запуска приложения", err)
	}
}
