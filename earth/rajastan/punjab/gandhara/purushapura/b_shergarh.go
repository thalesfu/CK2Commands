package purushapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室罗姞利呬ShergarhBarony struct {
	feud.BaseBarony
}

var BShergarh室罗姞利呬 feud.Barony = &室罗姞利呬ShergarhBarony{}

func init() {
    f := BShergarh室罗姞利呬.(*室罗姞利呬ShergarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shergarh",
		TitleName: "室罗姞利呬",
		TitleCode: "b_shergarh",
	}
}
