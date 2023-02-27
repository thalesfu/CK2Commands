package teluch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日耳曼尼亚GermaniasBarony struct {
	feud.BaseBarony
}

var BGermanias日耳曼尼亚 feud.Barony = &日耳曼尼亚GermaniasBarony{}

func init() {
    f := BGermanias日耳曼尼亚.(*日耳曼尼亚GermaniasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "germanias",
		TitleName: "日耳曼尼亚",
		TitleCode: "b_germanias",
	}
}
