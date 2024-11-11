package logic

import (
	"Friends/src/assets"
	"Friends/src/components/structures"
	"Friends/src/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mymmrac/telego"
)

func SearchGroupUID(bot *telego.Bot, msg telego.Message, user *structures.User) string {
	file, err := os.Open("D:/study/BMSTU/paradigms_structures_of_pl/IT-s_Friends/src/assets/db/structure.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return "-1"
	}
	defer file.Close()

	fmt.Println("Я прочел файл")
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return "-1"
	}

	var result structures.Final
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return "-1"
	}

	filial_index := utils.IndexOf(assets.GetFilials(), user.Filial)
	course_index := utils.IndexOf(assets.GetCourses(user.Filial, user.Faculty, user.Cathedra), user.Course)
	faculty_index := utils.IndexOf(assets.GetFaculties(user.Filial), user.Faculty)
	cathedra_index := utils.IndexOf(assets.GetCathedras(user.Filial, user.Faculty), user.Cathedra)
	group_index := utils.IndexOf(assets.GetGroups(user.Filial, user.Course, user.Faculty, user.Cathedra), user.Group)

	// fmt.Println("Я нашел индекс филлиалаааааа ", filial_index)
	// fmt.Println("Я нашел индекс курса ", course_index)
	// fmt.Println("Я нашел индекс факультета ", faculty_index)
	// fmt.Println("Я нашел индекс кафедры ", cathedra_index)
	// fmt.Println("Я нашел индекс группы ", group_index)

	fmt.Println("UUUUIID:", result.Data.Children[filial_index].Children[faculty_index].Children[cathedra_index].Children[course_index].Children[group_index].Uuid)
	return result.Data.Children[filial_index].Children[faculty_index].Children[cathedra_index].Children[course_index].Children[group_index].Uuid
}
