package yopurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 范家小掌FanjiaxiaozhangBarony struct {
	feud.BaseBarony
}

var BFanjiaxiaozhang范家小掌 feud.Barony = &范家小掌FanjiaxiaozhangBarony{}

func init() {
	f := BFanjiaxiaozhang范家小掌.(*范家小掌FanjiaxiaozhangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fanjiaxiaozhang",
		TitleName: "范家小掌",
		TitleCode: "b_fanjiaxiaozhang",
	}
}
