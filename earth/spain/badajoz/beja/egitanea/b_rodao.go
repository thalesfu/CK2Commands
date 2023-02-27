package egitanea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗当RodaoBarony struct {
	feud.BaseBarony
}

var BRodao罗当 feud.Barony = &罗当RodaoBarony{}

func init() {
    f := BRodao罗当.(*罗当RodaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rodao",
		TitleName: "罗当",
		TitleCode: "b_rodao",
	}
}
