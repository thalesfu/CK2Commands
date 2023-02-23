package valabhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩地耶MadyaBarony struct {
	feud.BaseBarony
}

var BMadya摩地耶 feud.Barony = &摩地耶MadyaBarony{}

func init() {
	f := BMadya摩地耶.(*摩地耶MadyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madya",
		TitleName: "摩地耶",
		TitleCode: "b_madya",
	}
}
