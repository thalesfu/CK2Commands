package kola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马费尔斯科伊MafelskoiBarony struct {
	feud.BaseBarony
}

var BMafelskoi马费尔斯科伊 feud.Barony = &马费尔斯科伊MafelskoiBarony{}

func init() {
	f := BMafelskoi马费尔斯科伊.(*马费尔斯科伊MafelskoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mafelskoi",
		TitleName: "马费尔斯科伊",
		TitleCode: "b_mafelskoi",
	}
}
