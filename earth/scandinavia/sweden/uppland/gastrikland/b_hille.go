package gastrikland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希勒HilleBarony struct {
	feud.BaseBarony
}

var BHille希勒 feud.Barony = &希勒HilleBarony{}

func init() {
	f := BHille希勒.(*希勒HilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hille",
		TitleName: "希勒",
		TitleCode: "b_hille",
	}
}
