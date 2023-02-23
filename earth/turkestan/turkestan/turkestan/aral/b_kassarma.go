package aral

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡萨尔马KassarmaBarony struct {
	feud.BaseBarony
}

var BKassarma卡萨尔马 feud.Barony = &卡萨尔马KassarmaBarony{}

func init() {
	f := BKassarma卡萨尔马.(*卡萨尔马KassarmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kassarma",
		TitleName: "卡萨尔马",
		TitleCode: "b_kassarma",
	}
}
