package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[255] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[255][255] = People_255
	HistoryPeopleMap[255][40255] = People_40255
	HistoryPeopleMap[255][70255] = People_70255
	HistoryPeopleMap[255][160255] = People_160255
	HistoryPeopleMap[255][170255] = People_170255
	HistoryPeopleMap[255][180255] = People_180255
	HistoryPeopleMap[255][190255] = People_190255
	HistoryPeopleMap[255][200255] = People_200255
	HistoryPeopleMap[255][260255] = People_260255
}
