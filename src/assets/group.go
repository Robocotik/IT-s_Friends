package assets

import (
	"Friends/src/components/structures"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Group [11]string = [11]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}

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

	var result structures.Final
	_ = json.Unmarshal(data, &result)
	groups = []string{}
	for _, group := range result.Data.Children {
		groups = append(groups, group.Abbr)
	}
	fmt.Sprintf("ГРУППЫ: ", groups)
	return groups
}
