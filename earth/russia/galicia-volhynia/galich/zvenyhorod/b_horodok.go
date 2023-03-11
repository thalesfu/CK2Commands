package zvenyhorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈罗多克HorodokBarony struct {
	feud.BaseBarony
}

var BHorodok戈罗多克 feud.Barony = &戈罗多克HorodokBarony{}

func init() {
    f := BHorodok戈罗多克.(*戈罗多克HorodokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "horodok",
		TitleName: "戈罗多克",
		TitleCode: "b_horodok",
	}
}
