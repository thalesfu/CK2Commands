package perigord

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尚瑟拉德ChanceladeBarony struct {
	feud.BaseBarony
}

var BChancelade尚瑟拉德 feud.Barony = &尚瑟拉德ChanceladeBarony{}

func init() {
	f := BChancelade尚瑟拉德.(*尚瑟拉德ChanceladeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chancelade",
		TitleName: "尚瑟拉德",
		TitleCode: "b_chancelade",
	}
}
