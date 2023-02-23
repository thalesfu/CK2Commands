package ancona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥西莫OsimoBarony struct {
	feud.BaseBarony
}

var BOsimo奥西莫 feud.Barony = &奥西莫OsimoBarony{}

func init() {
	f := BOsimo奥西莫.(*奥西莫OsimoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osimo",
		TitleName: "奥西莫",
		TitleCode: "b_osimo",
	}
}
