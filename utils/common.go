package utils

import (
	"fmt"
	"os"
)

// TakeCurrentDirectory получает путь до текущей директории
func TakeCurrentDirectory() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

// // writeInFileMap записывает данные в файл что-бы видеть что находится в мапе
// func (d *Data) writeInFileMap() {
// 	file, err := os.Create("text.txt")
// 	if err != nil {
// 		log.Println("file open error")
// 	}
// 	defer func(file *os.File) {
// 		err := file.Close()
// 		if err != nil {
// 			log.Println("error when closing the file")
// 		}
// 	}(file)
// 	for i, item := range d.Mapa {
// 		_, err = file.Write([]byte(fmt.Sprintf("%s%v\n", i, item)))
// 		if err != nil {
// 			log.Println("file write error")
// 		}
// 	}
// }
