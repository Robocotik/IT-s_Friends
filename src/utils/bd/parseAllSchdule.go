package bd

import (
	"Friends/src/entities"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

func ParseAllSchdule(conn *pgx.Conn) error {
	resp, err := http.Get("https://lks.bmstu.ru/lks-back/api/v1/structure")
	if err != nil {
		fmt.Println("Ошибка при get запросе :", err)
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении DATA:", err)
		return err
	}
	var result entities.Final
	err = json.Unmarshal(data, &result)
	// fmt.Println("result +", result)
	for _, filial := range result.Data.Children {
		// fmt.Println("я нашел филлиал", filial.Abbr)
		for _, faculty := range filial.Children {
			// fmt.Println("я нашел факультет", filial.Abbr)
			for _, cathedra := range faculty.Children {
				// fmt.Println("я нашел кафедру", cathedra.Abbr)
				for _, course := range cathedra.Children {
					// fmt.Println("я нашел курс", course.Abbr)
					for _, group := range course.Children {
						// fmt.Println("я нашел группу", group.Abbr)
						_, err := conn.Exec(context.Background(),
							"INSERT INTO schedule (uuid, filial_title, faculty_title, course_title, cathedra_title, group_title) VALUES($1, $2, $3, $4, $5, $6) on conflict (uuid) do nothing",
							group.Uuid, filial.Abbr, faculty.Abbr, course.Abbr, cathedra.Abbr, group.Abbr)
						if err != nil {
							fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
							return err
						}

					}
				}
			}
		}
	}
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return err
	}
	return nil
}
