package utils

import (
	// "Friends/src/assets"
	"Friends/src/components/structures"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func SearchGroupUID(bot *telego.Bot, msg telego.Message, filial string, course string, faculty string, cathedra string) {
	file, err := os.Open("D:/study/BMSTU/paradigms_structures_of_pl/IT-s_Friends/src/assets/db/structure.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()
	_, _ = bot.SendMessage(tu.Message(
		msg.Chat.ChatID(),
		fmt.Sprintf("Я прочел файл"),
	))
	fmt.Println("Я прочел файл")
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

	// filial_index := -1
	// for index, filial_asset := range assets.Filials {
	// 	if filial == filial_asset {
	// 		filial_index = index
	// 		break
	// 	}
	// }
	// fmt.Println("Я нашел индекс филлиала ", filial_index)

	// if filial_index == -1 {
	// 	fmt.Println("Филиал не найден")
	// 	return
	// }

	for _, faculty := range result.Data.Children[0].Children {
		for _, group := range faculty.Children {
			_, _ = bot.SendMessage(tu.Message(
				msg.Chat.ChatID(),
				fmt.Sprintf("Group: ", group.Abbr),
			))
		}
	}
}
