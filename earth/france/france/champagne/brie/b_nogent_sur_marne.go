package brie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺让Nogent_sur_marneBarony struct {
	feud.BaseBarony
}

var BNogent_sur_marne诺让 feud.Barony = &诺让Nogent_sur_marneBarony{}

func init() {
    f := BNogent_sur_marne诺让.(*诺让Nogent_sur_marneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nogent_sur_marne",
		TitleName: "诺让",
		TitleCode: "b_nogent_sur_marne",
	}
}
