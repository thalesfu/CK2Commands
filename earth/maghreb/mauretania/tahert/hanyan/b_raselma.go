package hanyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉斯马RaselmaBarony struct {
	feud.BaseBarony
}

var BRaselma拉斯马 feud.Barony = &拉斯马RaselmaBarony{}

func init() {
	f := BRaselma拉斯马.(*拉斯马RaselmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raselma",
		TitleName: "拉斯马",
		TitleCode: "b_raselma",
	}
}
