package fes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因吉勒EnjilBarony struct {
	feud.BaseBarony
}

var BEnjil因吉勒 feud.Barony = &因吉勒EnjilBarony{}

func init() {
    f := BEnjil因吉勒.(*因吉勒EnjilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "enjil",
		TitleName: "因吉勒",
		TitleCode: "b_enjil",
	}
}
