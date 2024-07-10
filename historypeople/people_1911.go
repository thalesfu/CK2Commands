package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1911] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1911][71911] = People_71911
	HistoryPeopleMap[1911][81911] = People_81911
	HistoryPeopleMap[1911][161911] = People_161911
	HistoryPeopleMap[1911][191911] = People_191911
	HistoryPeopleMap[1911][461911] = People_461911
}
