package kuban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别什陶BeshtauBarony struct {
	feud.BaseBarony
}

var BBeshtau别什陶 feud.Barony = &别什陶BeshtauBarony{}

func init() {
    f := BBeshtau别什陶.(*别什陶BeshtauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beshtau",
		TitleName: "别什陶",
		TitleCode: "b_beshtau",
	}
}
