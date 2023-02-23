package evora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维迪堡CastelodevideBarony struct {
	feud.BaseBarony
}

var BCastelodevide维迪堡 feud.Barony = &维迪堡CastelodevideBarony{}

func init() {
	f := BCastelodevide维迪堡.(*维迪堡CastelodevideBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castelodevide",
		TitleName: "维迪堡",
		TitleCode: "b_castelodevide",
	}
}
