package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7255] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7255][7255] = People_7255
	HistoryPeopleMap[7255][167255] = People_167255
	HistoryPeopleMap[7255][217255] = People_217255
	HistoryPeopleMap[7255][247255] = People_247255
}
