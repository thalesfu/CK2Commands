package frankfurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法兰克福FrankfurtBarony struct {
	feud.BaseBarony
}

var BFrankfurt法兰克福 feud.Barony = &法兰克福FrankfurtBarony{}

func init() {
	f := BFrankfurt法兰克福.(*法兰克福FrankfurtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "frankfurt",
		TitleName: "法兰克福",
		TitleCode: "b_frankfurt",
	}
}
