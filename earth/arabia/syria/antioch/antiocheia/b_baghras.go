package antiocheia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴格拉斯BaghrasBarony struct {
	feud.BaseBarony
}

var BBaghras巴格拉斯 feud.Barony = &巴格拉斯BaghrasBarony{}

func init() {
	f := BBaghras巴格拉斯.(*巴格拉斯BaghrasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baghras",
		TitleName: "巴格拉斯",
		TitleCode: "b_baghras",
	}
}
