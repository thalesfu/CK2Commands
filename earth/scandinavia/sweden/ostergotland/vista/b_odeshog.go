package vista

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄德斯赫格OdeshogBarony struct {
	feud.BaseBarony
}

var BOdeshog厄德斯赫格 feud.Barony = &厄德斯赫格OdeshogBarony{}

func init() {
	f := BOdeshog厄德斯赫格.(*厄德斯赫格OdeshogBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "odeshog",
		TitleName: "厄德斯赫格",
		TitleCode: "b_odeshog",
	}
}
