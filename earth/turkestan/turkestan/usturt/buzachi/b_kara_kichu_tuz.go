package buzachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉_基丘_图兹Kara_kichu_tuzBarony struct {
	feud.BaseBarony
}

var BKara_kichu_tuz卡拉_基丘_图兹 feud.Barony = &卡拉_基丘_图兹Kara_kichu_tuzBarony{}

func init() {
    f := BKara_kichu_tuz卡拉_基丘_图兹.(*卡拉_基丘_图兹Kara_kichu_tuzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kara_kichu_tuz",
		TitleName: "卡拉_基丘_图兹",
		TitleCode: "b_kara_kichu_tuz",
	}
}
