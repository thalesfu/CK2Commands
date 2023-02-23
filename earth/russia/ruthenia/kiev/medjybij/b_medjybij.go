package medjybij

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅德日比日MedjybijBarony struct {
	feud.BaseBarony
}

var BMedjybij梅德日比日 feud.Barony = &梅德日比日MedjybijBarony{}

func init() {
	f := BMedjybij梅德日比日.(*梅德日比日MedjybijBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medjybij",
		TitleName: "梅德日比日",
		TitleCode: "b_medjybij",
	}
}
