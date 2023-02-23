package amorion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿摩里翁AmorionBarony struct {
	feud.BaseBarony
}

var BAmorion阿摩里翁 feud.Barony = &阿摩里翁AmorionBarony{}

func init() {
	f := BAmorion阿摩里翁.(*阿摩里翁AmorionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amorion",
		TitleName: "阿摩里翁",
		TitleCode: "b_amorion",
	}
}
