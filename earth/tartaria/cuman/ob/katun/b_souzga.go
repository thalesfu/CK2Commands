package katun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索乌兹加SouzgaBarony struct {
	feud.BaseBarony
}

var BSouzga索乌兹加 feud.Barony = &索乌兹加SouzgaBarony{}

func init() {
    f := BSouzga索乌兹加.(*索乌兹加SouzgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "souzga",
		TitleName: "索乌兹加",
		TitleCode: "b_souzga",
	}
}
