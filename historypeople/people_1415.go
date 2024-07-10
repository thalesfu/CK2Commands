package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1415] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1415][1415] = People_1415
	HistoryPeopleMap[1415][31415] = People_31415
	HistoryPeopleMap[1415][71415] = People_71415
	HistoryPeopleMap[1415][91415] = People_91415
	HistoryPeopleMap[1415][161415] = People_161415
	HistoryPeopleMap[1415][191415] = People_191415
}
