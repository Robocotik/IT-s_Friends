package utils

import (
	"Friends/src/components/structures"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func SearchGroupUID(filal string, course string, faculty string, cathedra string) {

	file, err := os.Open("D:/study/BMSTU/paradigms_structures_of_pl/IT-s_Friends/src/assets/db/structure.json")
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

	var result structures.Final
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}

	for _, faculty := range result.Data.Children[0].Children {
		for _, group := range faculty.Children {
			fmt.Println("Group:", group.Abbr)
		}
	}

}
