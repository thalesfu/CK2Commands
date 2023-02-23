package lindisfarne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 林迪斯法恩LindisfarneBarony struct {
	feud.BaseBarony
}

var BLindisfarne林迪斯法恩 feud.Barony = &林迪斯法恩LindisfarneBarony{}

func init() {
	f := BLindisfarne林迪斯法恩.(*林迪斯法恩LindisfarneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lindisfarne",
		TitleName: "林迪斯法恩",
		TitleCode: "b_lindisfarne",
	}
}
