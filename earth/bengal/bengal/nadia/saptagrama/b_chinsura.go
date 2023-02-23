package saptagrama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 真苏罗ChinsuraBarony struct {
	feud.BaseBarony
}

var BChinsura真苏罗 feud.Barony = &真苏罗ChinsuraBarony{}

func init() {
	f := BChinsura真苏罗.(*真苏罗ChinsuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chinsura",
		TitleName: "真苏罗",
		TitleCode: "b_chinsura",
	}
}
