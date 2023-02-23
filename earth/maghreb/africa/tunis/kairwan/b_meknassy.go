package kairwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅克纳萨伊MeknassyBarony struct {
	feud.BaseBarony
}

var BMeknassy梅克纳萨伊 feud.Barony = &梅克纳萨伊MeknassyBarony{}

func init() {
	f := BMeknassy梅克纳萨伊.(*梅克纳萨伊MeknassyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meknassy",
		TitleName: "梅克纳萨伊",
		TitleCode: "b_meknassy",
	}
}
