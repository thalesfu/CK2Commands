package trinkitat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特卜TebBarony struct {
	feud.BaseBarony
}

var BTeb特卜 feud.Barony = &特卜TebBarony{}

func init() {
	f := BTeb特卜.(*特卜TebBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teb",
		TitleName: "特卜",
		TitleCode: "b_teb",
	}
}
