package assets

import (
	"Friends/src/assets/emoji"
	"Friends/src/components/structures"
	"Friends/src/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetCourses(filial string) []string {
	var courses []string
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
	courses = []string{}

	filial_index := utils.IndexOf(GetFilials(), filial)
	fmt.Println("Я нашел индекс филлиалааа ", filial_index)

	for index, course := range result.Data.Children[filial_index].Children {
		courses = append(courses, course.Abbr + emoji.Courses[index])
	}
	fmt.Sprintf("КУРСЫ: ", courses)
	return courses
}
