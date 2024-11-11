package assets

import (
	"Friends/src/components/structures"
	"Friends/src/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)


func GetFaculties(filial string) []string {
	var faculties []string
	file, err := os.Open("D:/study/BMSTU/paradigms_structures_of_pl/IT-s_Friends/src/assets/db/structure.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return nil
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return nil
	}

	var result structures.Final
	_ = json.Unmarshal(data, &result)
	faculties = []string{}

	filial_index := utils.IndexOf(GetFilials(), filial)
	fmt.Println("Я нашел индекс филлиалаааа ", filial_index)

	for _, faculty := range result.Data.Children[filial_index].Children {
		faculties = append(faculties, faculty.Abbr)
	}
	fmt.Sprintf("Факультеты: ", faculties)
	return faculties
}
