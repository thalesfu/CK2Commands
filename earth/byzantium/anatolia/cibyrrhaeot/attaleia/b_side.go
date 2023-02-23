package attaleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西德SideBarony struct {
	feud.BaseBarony
}

var BSide西德 feud.Barony = &西德SideBarony{}

func init() {
	f := BSide西德.(*西德SideBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "side",
		TitleName: "西德",
		TitleCode: "b_side",
	}
}
