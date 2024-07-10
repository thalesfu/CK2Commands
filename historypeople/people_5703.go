package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5703] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5703][125703] = People_125703
	HistoryPeopleMap[5703][145703] = People_145703
	HistoryPeopleMap[5703][205703] = People_205703
}
