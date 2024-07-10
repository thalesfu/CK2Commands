package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1080] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1080][31080] = People_31080
	HistoryPeopleMap[1080][71080] = People_71080
	HistoryPeopleMap[1080][91080] = People_91080
	HistoryPeopleMap[1080][161080] = People_161080
	HistoryPeopleMap[1080][191080] = People_191080
	HistoryPeopleMap[1080][261080] = People_261080
}
