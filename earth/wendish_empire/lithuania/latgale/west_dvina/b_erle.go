package west_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔勒ErleBarony struct {
	feud.BaseBarony
}

var BErle埃尔勒 feud.Barony = &埃尔勒ErleBarony{}

func init() {
    f := BErle埃尔勒.(*埃尔勒ErleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erle",
		TitleName: "埃尔勒",
		TitleCode: "b_erle",
	}
}
