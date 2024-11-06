package utils

import (
	"Friends/src/components/structures"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)


func SearchGroupUID() {

	file, err := os.Open("D:/study/BMSTU/paradigms_structures_of_pl/IT-s_Friends/src/assets/db/structure.json")
	// file, err := filepath.Abs("../mypackage/data/file.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	// Декодируем JSON
	var result structures.Final
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}

	// Обращаемся к полю children
	for _, fillial := range result.Data.Children { // Доступ к Fillial
    for _, faculty := range fillial.Children { // Доступ к Faculty
        for _, group := range faculty.Children { // Доступ к Group
            fmt.Println("Group:", group)
        }
    }
}
}
