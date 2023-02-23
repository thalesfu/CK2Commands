package katun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼热罗克ManzherokBarony struct {
	feud.BaseBarony
}

var BManzherok曼热罗克 feud.Barony = &曼热罗克ManzherokBarony{}

func init() {
	f := BManzherok曼热罗克.(*曼热罗克ManzherokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manzherok",
		TitleName: "曼热罗克",
		TitleCode: "b_manzherok",
	}
}
