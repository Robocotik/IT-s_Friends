package logic

import (
	"errors"

	entities "github.com/Robocotik/IT-s_Friends/internal/models/entities"
)

func GetTimetableIntarsection(meSchedule []entities.IDay, friendSchedule []entities.IDay, day int) (int, int, error) {
	first := 20
	second := -1

	if len(meSchedule) == 0 || len(friendSchedule) == 0 {
		return -1, -1, errors.New("schedule empty")
	}
	for _, meDay := range meSchedule {
		if meDay.Discipline.Abbr == "Самостоятельная работа" && meDay.Day == day {
			return -1, -1, errors.New("rest time")
		}
	}

	for _, friendDay := range friendSchedule {
		if friendDay.Discipline.Abbr == "Самостоятельная работа" && friendDay.Day == day {
			return -1, -1, errors.New("rest time")
		}
	}

	for _, friendDay := range friendSchedule {
		if friendDay.Day != day {
			continue
		}
		for _, meDay := range meSchedule {
			if friendDay.Day == day && day == meDay.Day {
				if friendDay.Time == meDay.Time {
					if friendDay.Time < first {
						first = friendDay.Time
					}
					if meDay.Time > second {
						second = meDay.Time
					}

				}
			}
		}
	}

	if (first == 20 && second == -1){
		return -1, -1, errors.New("no interaction")
	}



	return first, second, nil
}
