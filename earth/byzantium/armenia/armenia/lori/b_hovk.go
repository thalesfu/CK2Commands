package lori

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 何福HovkBarony struct {
	feud.BaseBarony
}

var BHovk何福 feud.Barony = &何福HovkBarony{}

func init() {
    f := BHovk何福.(*何福HovkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hovk",
		TitleName: "何福",
		TitleCode: "b_hovk",
	}
}
