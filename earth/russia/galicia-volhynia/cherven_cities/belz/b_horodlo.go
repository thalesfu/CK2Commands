package belz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍罗德沃HorodloBarony struct {
	feud.BaseBarony
}

var BHorodlo霍罗德沃 feud.Barony = &霍罗德沃HorodloBarony{}

func init() {
    f := BHorodlo霍罗德沃.(*霍罗德沃HorodloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "horodlo",
		TitleName: "霍罗德沃",
		TitleCode: "b_horodlo",
	}
}
