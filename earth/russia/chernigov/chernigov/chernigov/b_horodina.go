package chernigov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈罗迪纳HorodinaBarony struct {
	feud.BaseBarony
}

var BHorodina戈罗迪纳 feud.Barony = &戈罗迪纳HorodinaBarony{}

func init() {
    f := BHorodina戈罗迪纳.(*戈罗迪纳HorodinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "horodina",
		TitleName: "戈罗迪纳",
		TitleCode: "b_horodina",
	}
}
