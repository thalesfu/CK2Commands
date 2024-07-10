package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1601] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1601][41601] = People_41601
	HistoryPeopleMap[1601][71601] = People_71601
	HistoryPeopleMap[1601][131601] = People_131601
	HistoryPeopleMap[1601][161601] = People_161601
	HistoryPeopleMap[1601][191601] = People_191601
}
