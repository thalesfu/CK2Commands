package kastamon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯塔莫努KastamonBarony struct {
	feud.BaseBarony
}

var BKastamon卡斯塔莫努 feud.Barony = &卡斯塔莫努KastamonBarony{}

func init() {
    f := BKastamon卡斯塔莫努.(*卡斯塔莫努KastamonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kastamon",
		TitleName: "卡斯塔莫努",
		TitleCode: "b_kastamon",
	}
}
