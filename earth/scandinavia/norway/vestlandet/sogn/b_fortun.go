package sogn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福通FortunBarony struct {
	feud.BaseBarony
}

var BFortun福通 feud.Barony = &福通FortunBarony{}

func init() {
	f := BFortun福通.(*福通FortunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fortun",
		TitleName: "福通",
		TitleCode: "b_fortun",
	}
}
