package abakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别连扎克BerenzhakBarony struct {
	feud.BaseBarony
}

var BBerenzhak别连扎克 feud.Barony = &别连扎克BerenzhakBarony{}

func init() {
    f := BBerenzhak别连扎克.(*别连扎克BerenzhakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berenzhak",
		TitleName: "别连扎克",
		TitleCode: "b_berenzhak",
	}
}
