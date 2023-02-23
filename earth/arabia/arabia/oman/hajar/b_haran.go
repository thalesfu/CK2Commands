package hajar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈兰HaranBarony struct {
	feud.BaseBarony
}

var BHaran哈兰 feud.Barony = &哈兰HaranBarony{}

func init() {
	f := BHaran哈兰.(*哈兰HaranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haran",
		TitleName: "哈兰",
		TitleCode: "b_haran",
	}
}
