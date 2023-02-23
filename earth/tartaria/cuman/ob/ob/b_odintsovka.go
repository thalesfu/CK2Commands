package ob

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥金佐夫卡OdintsovkaBarony struct {
	feud.BaseBarony
}

var BOdintsovka奥金佐夫卡 feud.Barony = &奥金佐夫卡OdintsovkaBarony{}

func init() {
	f := BOdintsovka奥金佐夫卡.(*奥金佐夫卡OdintsovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "odintsovka",
		TitleName: "奥金佐夫卡",
		TitleCode: "b_odintsovka",
	}
}
