package wag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉舍瓦LashwaBarony struct {
	feud.BaseBarony
}

var BLashwa拉舍瓦 feud.Barony = &拉舍瓦LashwaBarony{}

func init() {
    f := BLashwa拉舍瓦.(*拉舍瓦LashwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lashwa",
		TitleName: "拉舍瓦",
		TitleCode: "b_lashwa",
	}
}
