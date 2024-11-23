package assets

import (
	"Friends/src/assets/emoji"
	"Friends/src/entities"
	"Friends/src/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func GetCourses(filial string, faculty string, cathedra string) []string {
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

	var result entities.Final
	_ = json.Unmarshal(data, &result)
	courses = []string{}

	filial_index := utils.IndexOf(GetFilials(), filial)
	// fmt.Println("Я нашел индекс филлиалааа1 ", filial_index)
	cathedra_index := utils.IndexOf(GetCathedras(filial, faculty), cathedra)
	// fmt.Println("Я нашел индекс кафедрыыыыы1 ", cathedra_index)
	faculty_index := utils.IndexOf(GetFaculties(filial), faculty)
	// fmt.Println("Я нашел индекс факультета1 ", faculty_index)

	for index, course := range result.Data.Children[filial_index].Children[faculty_index].Children[cathedra_index].Children {
		fmt.Println("КУРС: ", course.Course)
		courses = append(courses, strconv.Itoa(course.Course)+emoji.Courses[index])
	}
	fmt.Println("КУРСЫ: ", courses)
	return courses
}
