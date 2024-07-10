package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5556] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5556][205556] = People_205556
	HistoryPeopleMap[5556][455556] = People_455556
	HistoryPeopleMap[5556][475556] = People_475556
}
