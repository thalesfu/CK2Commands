package uppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡格蒂纳SigtunaBarony struct {
	feud.BaseBarony
}

var BSigtuna锡格蒂纳 feud.Barony = &锡格蒂纳SigtunaBarony{}

func init() {
    f := BSigtuna锡格蒂纳.(*锡格蒂纳SigtunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sigtuna",
		TitleName: "锡格蒂纳",
		TitleCode: "b_sigtuna",
	}
}
