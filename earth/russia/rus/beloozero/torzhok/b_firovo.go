package torzhok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲罗沃FirovoBarony struct {
	feud.BaseBarony
}

var BFirovo菲罗沃 feud.Barony = &菲罗沃FirovoBarony{}

func init() {
	f := BFirovo菲罗沃.(*菲罗沃FirovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "firovo",
		TitleName: "菲罗沃",
		TitleCode: "b_firovo",
	}
}
