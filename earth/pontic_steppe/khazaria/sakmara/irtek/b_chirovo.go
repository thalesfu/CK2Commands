package irtek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇罗沃ChirovoBarony struct {
	feud.BaseBarony
}

var BChirovo奇罗沃 feud.Barony = &奇罗沃ChirovoBarony{}

func init() {
    f := BChirovo奇罗沃.(*奇罗沃ChirovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chirovo",
		TitleName: "奇罗沃",
		TitleCode: "b_chirovo",
	}
}
