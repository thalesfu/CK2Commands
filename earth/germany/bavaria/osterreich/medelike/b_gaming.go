package medelike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加明GamingBarony struct {
	feud.BaseBarony
}

var BGaming加明 feud.Barony = &加明GamingBarony{}

func init() {
    f := BGaming加明.(*加明GamingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gaming",
		TitleName: "加明",
		TitleCode: "b_gaming",
	}
}
