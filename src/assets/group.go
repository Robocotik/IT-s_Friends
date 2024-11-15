package assets

import (
	"Friends/src/entities"
	"Friends/src/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetGroups(filial string, course string, faculty string, cathedra string) []string {
	var groups []string
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

	filial_index := utils.IndexOf(GetFilials(), filial)
	cathedra_index := utils.IndexOf(GetCathedras(filial, faculty), cathedra)
	faculty_index := utils.IndexOf(GetFaculties(filial), faculty)
	course_index := utils.IndexOf(GetCourses(filial, faculty, cathedra), course)

	var result entities.Final
	_ = json.Unmarshal(data, &result)
	groups = []string{}

	for _, group := range result.Data.Children[filial_index].Children[faculty_index].Children[cathedra_index].Children[course_index].Children {
		groups = append(groups, group.Abbr)
	}

	fmt.Sprintf("ГРУППЫ: ", groups)
	return groups
}
