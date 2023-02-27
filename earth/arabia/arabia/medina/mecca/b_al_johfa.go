package mecca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约法Al_johfaBarony struct {
	feud.BaseBarony
}

var BAl_johfa约法 feud.Barony = &约法Al_johfaBarony{}

func init() {
    f := BAl_johfa约法.(*约法Al_johfaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_johfa",
		TitleName: "约法",
		TitleCode: "b_al_johfa",
	}
}
