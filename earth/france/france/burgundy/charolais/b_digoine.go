package charolais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪关DigoineBarony struct {
	feud.BaseBarony
}

var BDigoine迪关 feud.Barony = &迪关DigoineBarony{}

func init() {
    f := BDigoine迪关.(*迪关DigoineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "digoine",
		TitleName: "迪关",
		TitleCode: "b_digoine",
	}
}
