package utils

import (
	"fmt"
	"log"
	"os"
	"path"
)

func OpenFile(file string) (*os.File, error) {
	filePath := path.Join("R:", "Thales", "Game", "SteamLibrary", "steamapps", "common", "Crusader Kings II", file+".txt")
	fmt.Println(filePath)

	return os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
}

func CloseFile(file *os.File) {
	func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)
}
