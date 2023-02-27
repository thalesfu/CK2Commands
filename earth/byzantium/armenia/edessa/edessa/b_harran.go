package edessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈兰HarranBarony struct {
	feud.BaseBarony
}

var BHarran哈兰 feud.Barony = &哈兰HarranBarony{}

func init() {
    f := BHarran哈兰.(*哈兰HarranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harran",
		TitleName: "哈兰",
		TitleCode: "b_harran",
	}
}
