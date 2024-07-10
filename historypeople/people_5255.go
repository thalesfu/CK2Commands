package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5255] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5255][125255] = People_125255
	HistoryPeopleMap[5255][145255] = People_145255
	HistoryPeopleMap[5255][155255] = People_155255
	HistoryPeopleMap[5255][205255] = People_205255
}
