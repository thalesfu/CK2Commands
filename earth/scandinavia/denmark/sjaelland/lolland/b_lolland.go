package lolland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛兰LollandBarony struct {
	feud.BaseBarony
}

var BLolland洛兰 feud.Barony = &洛兰LollandBarony{}

func init() {
	f := BLolland洛兰.(*洛兰LollandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lolland",
		TitleName: "洛兰",
		TitleCode: "b_lolland",
	}
}
