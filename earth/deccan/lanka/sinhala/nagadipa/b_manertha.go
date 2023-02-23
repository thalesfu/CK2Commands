package nagadipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩涅他ManerthaBarony struct {
	feud.BaseBarony
}

var BManertha摩涅他 feud.Barony = &摩涅他ManerthaBarony{}

func init() {
	f := BManertha摩涅他.(*摩涅他ManerthaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manertha",
		TitleName: "摩涅他",
		TitleCode: "b_manertha",
	}
}
