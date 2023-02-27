package isle_of_man

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮尔PeelBarony struct {
	feud.BaseBarony
}

var BPeel皮尔 feud.Barony = &皮尔PeelBarony{}

func init() {
    f := BPeel皮尔.(*皮尔PeelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "peel",
		TitleName: "皮尔",
		TitleCode: "b_peel",
	}
}
