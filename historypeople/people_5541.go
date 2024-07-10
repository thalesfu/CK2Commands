package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5541] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5541][205541] = People_205541
	HistoryPeopleMap[5541][455541] = People_455541
}
