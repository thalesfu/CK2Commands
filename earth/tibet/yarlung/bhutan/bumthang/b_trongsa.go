package bumthang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 通萨TrongsaBarony struct {
	feud.BaseBarony
}

var BTrongsa通萨 feud.Barony = &通萨TrongsaBarony{}

func init() {
	f := BTrongsa通萨.(*通萨TrongsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trongsa",
		TitleName: "通萨",
		TitleCode: "b_trongsa",
	}
}
