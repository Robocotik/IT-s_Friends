package logic

import (
	"Friends/src/assets"
	"Friends/src/components/structures"
	"Friends/src/entities"
	"Friends/src/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mymmrac/telego"
)

func SearchGroupUID(bot *telego.Bot, msg telego.Message, friend *structures.AskedFriend) string {
	file, err := os.Open("D:/study/BMSTU/paradigms_structures_of_pl/IT-s_Friends/src/assets/db/structure.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return "-1"
	}
	defer file.Close()

	// fmt.Println("Я прочел файл")
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return "-1"
	}

	var result entities.Final
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return "-1"
	}

	filial_index := utils.IndexOf(assets.GetFilials(), friend.Filial) // Переписать 
	course_index := utils.IndexOf(assets.GetCourses(friend.Filial, friend.Faculty, friend.Cathedra), friend.Course)
	faculty_index := utils.IndexOf(assets.GetFaculties(friend.Filial), friend.Faculty)
	cathedra_index := utils.IndexOf(assets.GetCathedras(friend.Filial, friend.Faculty), friend.Cathedra)
	group_index := utils.IndexOf(assets.GetGroups(friend.Filial, friend.Course, friend.Faculty, friend.Cathedra), friend.Group)


	fmt.Println("UUUUIID:", result.Data.Children[filial_index].Children[faculty_index].Children[cathedra_index].Children[course_index].Children[group_index].Uuid)
	return result.Data.Children[filial_index].Children[faculty_index].Children[cathedra_index].Children[course_index].Children[group_index].Uuid
}
