package mordva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因萨尔InsarBarony struct {
	feud.BaseBarony
}

var BInsar因萨尔 feud.Barony = &因萨尔InsarBarony{}

func init() {
    f := BInsar因萨尔.(*因萨尔InsarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "insar",
		TitleName: "因萨尔",
		TitleCode: "b_insar",
	}
}
