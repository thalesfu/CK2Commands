package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5101] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5101][45101] = People_45101
	HistoryPeopleMap[5101][125101] = People_125101
	HistoryPeopleMap[5101][145101] = People_145101
	HistoryPeopleMap[5101][155101] = People_155101
	HistoryPeopleMap[5101][205101] = People_205101
	HistoryPeopleMap[5101][235101] = People_235101
}
