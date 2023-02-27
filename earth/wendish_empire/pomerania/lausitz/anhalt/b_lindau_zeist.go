package anhalt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 林道Lindau_zeistBarony struct {
	feud.BaseBarony
}

var BLindau_zeist林道 feud.Barony = &林道Lindau_zeistBarony{}

func init() {
    f := BLindau_zeist林道.(*林道Lindau_zeistBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lindau_zeist",
		TitleName: "林道",
		TitleCode: "b_lindau_zeist",
	}
}
