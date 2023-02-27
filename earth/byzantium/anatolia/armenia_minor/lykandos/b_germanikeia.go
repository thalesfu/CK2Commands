package lykandos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日耳曼尼基亚GermanikeiaBarony struct {
	feud.BaseBarony
}

var BGermanikeia日耳曼尼基亚 feud.Barony = &日耳曼尼基亚GermanikeiaBarony{}

func init() {
    f := BGermanikeia日耳曼尼基亚.(*日耳曼尼基亚GermanikeiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "germanikeia",
		TitleName: "日耳曼尼基亚",
		TitleCode: "b_germanikeia",
	}
}
