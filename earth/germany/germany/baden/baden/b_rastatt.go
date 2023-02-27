package baden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉施塔特RastattBarony struct {
	feud.BaseBarony
}

var BRastatt拉施塔特 feud.Barony = &拉施塔特RastattBarony{}

func init() {
    f := BRastatt拉施塔特.(*拉施塔特RastattBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rastatt",
		TitleName: "拉施塔特",
		TitleCode: "b_rastatt",
	}
}
