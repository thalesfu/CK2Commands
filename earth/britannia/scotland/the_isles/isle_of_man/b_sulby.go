package isle_of_man

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔比SulbyBarony struct {
	feud.BaseBarony
}

var BSulby萨尔比 feud.Barony = &萨尔比SulbyBarony{}

func init() {
    f := BSulby萨尔比.(*萨尔比SulbyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sulby",
		TitleName: "萨尔比",
		TitleCode: "b_sulby",
	}
}
