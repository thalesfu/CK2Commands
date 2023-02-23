package apulia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉韦洛LavelloBarony struct {
	feud.BaseBarony
}

var BLavello拉韦洛 feud.Barony = &拉韦洛LavelloBarony{}

func init() {
	f := BLavello拉韦洛.(*拉韦洛LavelloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lavello",
		TitleName: "拉韦洛",
		TitleCode: "b_lavello",
	}
}
