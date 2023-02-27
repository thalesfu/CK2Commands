package belaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希波沃ShipovoBarony struct {
	feud.BaseBarony
}

var BShipovo希波沃 feud.Barony = &希波沃ShipovoBarony{}

func init() {
    f := BShipovo希波沃.(*希波沃ShipovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shipovo",
		TitleName: "希波沃",
		TitleCode: "b_shipovo",
	}
}
