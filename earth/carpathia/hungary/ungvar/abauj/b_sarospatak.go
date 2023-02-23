package abauj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙罗什保陶克SarospatakBarony struct {
	feud.BaseBarony
}

var BSarospatak沙罗什保陶克 feud.Barony = &沙罗什保陶克SarospatakBarony{}

func init() {
	f := BSarospatak沙罗什保陶克.(*沙罗什保陶克SarospatakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarospatak",
		TitleName: "沙罗什保陶克",
		TitleCode: "b_sarospatak",
	}
}
