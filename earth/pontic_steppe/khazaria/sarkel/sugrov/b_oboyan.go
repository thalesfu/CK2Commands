package sugrov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥博扬OboyanBarony struct {
	feud.BaseBarony
}

var BOboyan奥博扬 feud.Barony = &奥博扬OboyanBarony{}

func init() {
    f := BOboyan奥博扬.(*奥博扬OboyanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oboyan",
		TitleName: "奥博扬",
		TitleCode: "b_oboyan",
	}
}
