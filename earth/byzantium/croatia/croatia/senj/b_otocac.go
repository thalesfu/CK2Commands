package senj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥托查茨OtocacBarony struct {
	feud.BaseBarony
}

var BOtocac奥托查茨 feud.Barony = &奥托查茨OtocacBarony{}

func init() {
    f := BOtocac奥托查茨.(*奥托查茨OtocacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "otocac",
		TitleName: "奥托查茨",
		TitleCode: "b_otocac",
	}
}
