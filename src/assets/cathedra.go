package assets

import (
	"Friends/src/entities"
	"Friends/src/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetCathedras(filial string, faculty string) []string {
	var cathedras []string
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

	var result entities.Final
	_ = json.Unmarshal(data, &result)
	cathedras = []string{}

	filial_index := utils.IndexOf(GetFilials(), filial)

	fmt.Println("Я нашел индекс филлиалаа ", filial_index)

	faculty_index := utils.IndexOf(GetFaculties(filial), faculty)

	fmt.Println("Я нашел индекс факультета ", faculty_index)

	for _, cathedra := range result.Data.Children[filial_index].Children[faculty_index].Children {
		cathedras = append(cathedras, cathedra.Abbr)
	}
	fmt.Sprintf("Кафедры: ", cathedras)
	return cathedras
}
