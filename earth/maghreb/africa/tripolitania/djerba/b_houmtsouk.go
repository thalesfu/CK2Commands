package djerba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 豪迈特苏克HoumtsoukBarony struct {
	feud.BaseBarony
}

var BHoumtsouk豪迈特苏克 feud.Barony = &豪迈特苏克HoumtsoukBarony{}

func init() {
	f := BHoumtsouk豪迈特苏克.(*豪迈特苏克HoumtsoukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "houmtsouk",
		TitleName: "豪迈特苏克",
		TitleCode: "b_houmtsouk",
	}
}
