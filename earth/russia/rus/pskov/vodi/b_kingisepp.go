package vodi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚马KingiseppBarony struct {
	feud.BaseBarony
}

var BKingisepp亚马 feud.Barony = &亚马KingiseppBarony{}

func init() {
	f := BKingisepp亚马.(*亚马KingiseppBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kingisepp",
		TitleName: "亚马",
		TitleCode: "b_kingisepp",
	}
}
