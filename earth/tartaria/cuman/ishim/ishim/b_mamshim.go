package ishim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马姆希姆MamshimBarony struct {
	feud.BaseBarony
}

var BMamshim马姆希姆 feud.Barony = &马姆希姆MamshimBarony{}

func init() {
	f := BMamshim马姆希姆.(*马姆希姆MamshimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mamshim",
		TitleName: "马姆希姆",
		TitleCode: "b_mamshim",
	}
}
