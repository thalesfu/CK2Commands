package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2801] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2801][2801] = People_2801
	HistoryPeopleMap[2801][32801] = People_32801
	HistoryPeopleMap[2801][72801] = People_72801
	HistoryPeopleMap[2801][142801] = People_142801
	HistoryPeopleMap[2801][212801] = People_212801
	HistoryPeopleMap[2801][222801] = People_222801
	HistoryPeopleMap[2801][462801] = People_462801
	HistoryPeopleMap[2801][472801] = People_472801
}
