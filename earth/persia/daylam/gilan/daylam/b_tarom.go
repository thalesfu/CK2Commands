package daylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔罗姆TaromBarony struct {
	feud.BaseBarony
}

var BTarom塔罗姆 feud.Barony = &塔罗姆TaromBarony{}

func init() {
    f := BTarom塔罗姆.(*塔罗姆TaromBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarom",
		TitleName: "塔罗姆",
		TitleCode: "b_tarom",
	}
}
