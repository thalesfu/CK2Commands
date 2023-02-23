package yatenga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜巴雷DubareBarony struct {
	feud.BaseBarony
}

var BDubare杜巴雷 feud.Barony = &杜巴雷DubareBarony{}

func init() {
	f := BDubare杜巴雷.(*杜巴雷DubareBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dubare",
		TitleName: "杜巴雷",
		TitleCode: "b_dubare",
	}
}
