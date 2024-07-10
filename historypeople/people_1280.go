package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1280] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1280][1280] = People_1280
	HistoryPeopleMap[1280][71280] = People_71280
	HistoryPeopleMap[1280][91280] = People_91280
	HistoryPeopleMap[1280][161280] = People_161280
	HistoryPeopleMap[1280][191280] = People_191280
	HistoryPeopleMap[1280][261280] = People_261280
}
