package coqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 珠龙ZhuglungBarony struct {
	feud.BaseBarony
}

var BZhuglung珠龙 feud.Barony = &珠龙ZhuglungBarony{}

func init() {
	f := BZhuglung珠龙.(*珠龙ZhuglungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhuglung",
		TitleName: "珠龙",
		TitleCode: "b_zhuglung",
	}
}
