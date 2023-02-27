package zahedan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达姆恩HamunBarony struct {
	feud.BaseBarony
}

var BHamun达姆恩 feud.Barony = &达姆恩HamunBarony{}

func init() {
    f := BHamun达姆恩.(*达姆恩HamunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hamun",
		TitleName: "达姆恩",
		TitleCode: "b_hamun",
	}
}
