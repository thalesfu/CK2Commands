package salamanca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉曼卡SalamancaBarony struct {
	feud.BaseBarony
}

var BSalamanca萨拉曼卡 feud.Barony = &萨拉曼卡SalamancaBarony{}

func init() {
	f := BSalamanca萨拉曼卡.(*萨拉曼卡SalamancaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salamanca",
		TitleName: "萨拉曼卡",
		TitleCode: "b_salamanca",
	}
}
