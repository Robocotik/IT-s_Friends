package assets

import (
	"Friends/src/components/structures"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetFilials() []string {
	var filials []string
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
	filials = []string{}
	for _, filial := range result.Data.Children {
		filials = append(filials, filial.Abbr)
	}
	fmt.Sprintf("ФИЛЛИАЛЫ: ", filials)
	return filials
}
