package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1890] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1890][31890] = People_31890
	HistoryPeopleMap[1890][71890] = People_71890
	HistoryPeopleMap[1890][191890] = People_191890
}
