package hainaut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙勒罗瓦CharleroiBarony struct {
	feud.BaseBarony
}

var BCharleroi沙勒罗瓦 feud.Barony = &沙勒罗瓦CharleroiBarony{}

func init() {
	f := BCharleroi沙勒罗瓦.(*沙勒罗瓦CharleroiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charleroi",
		TitleName: "沙勒罗瓦",
		TitleCode: "b_charleroi",
	}
}
