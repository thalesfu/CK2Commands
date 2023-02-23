package derby

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯尔顿CastletonBarony struct {
	feud.BaseBarony
}

var BCastleton卡斯尔顿 feud.Barony = &卡斯尔顿CastletonBarony{}

func init() {
	f := BCastleton卡斯尔顿.(*卡斯尔顿CastletonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castleton",
		TitleName: "卡斯尔顿",
		TitleCode: "b_castleton",
	}
}
