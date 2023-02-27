package tsambagarav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尚塔TashantaBarony struct {
	feud.BaseBarony
}

var BTashanta塔尚塔 feud.Barony = &塔尚塔TashantaBarony{}

func init() {
    f := BTashanta塔尚塔.(*塔尚塔TashantaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tashanta",
		TitleName: "塔尚塔",
		TitleCode: "b_tashanta",
	}
}
