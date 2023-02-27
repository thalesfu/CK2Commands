package wolgast

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃格辛EggesinBarony struct {
	feud.BaseBarony
}

var BEggesin埃格辛 feud.Barony = &埃格辛EggesinBarony{}

func init() {
    f := BEggesin埃格辛.(*埃格辛EggesinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eggesin",
		TitleName: "埃格辛",
		TitleCode: "b_eggesin",
	}
}
