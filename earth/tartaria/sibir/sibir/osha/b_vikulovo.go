package osha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维库洛沃VikulovoBarony struct {
	feud.BaseBarony
}

var BVikulovo维库洛沃 feud.Barony = &维库洛沃VikulovoBarony{}

func init() {
    f := BVikulovo维库洛沃.(*维库洛沃VikulovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vikulovo",
		TitleName: "维库洛沃",
		TitleCode: "b_vikulovo",
	}
}
