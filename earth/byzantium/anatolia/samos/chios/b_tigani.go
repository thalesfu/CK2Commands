package chios

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提加尼TiganiBarony struct {
	feud.BaseBarony
}

var BTigani提加尼 feud.Barony = &提加尼TiganiBarony{}

func init() {
    f := BTigani提加尼.(*提加尼TiganiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tigani",
		TitleName: "提加尼",
		TitleCode: "b_tigani",
	}
}
