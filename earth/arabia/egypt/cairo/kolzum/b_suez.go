package kolzum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏伊士SuezBarony struct {
	feud.BaseBarony
}

var BSuez苏伊士 feud.Barony = &苏伊士SuezBarony{}

func init() {
    f := BSuez苏伊士.(*苏伊士SuezBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suez",
		TitleName: "苏伊士",
		TitleCode: "b_suez",
	}
}
