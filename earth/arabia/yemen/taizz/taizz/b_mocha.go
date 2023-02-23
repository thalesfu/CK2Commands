package taizz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩卡MochaBarony struct {
	feud.BaseBarony
}

var BMocha摩卡 feud.Barony = &摩卡MochaBarony{}

func init() {
	f := BMocha摩卡.(*摩卡MochaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mocha",
		TitleName: "摩卡",
		TitleCode: "b_mocha",
	}
}
