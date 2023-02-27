package nagauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗努恩LadnunBarony struct {
	feud.BaseBarony
}

var BLadnun罗努恩 feud.Barony = &罗努恩LadnunBarony{}

func init() {
    f := BLadnun罗努恩.(*罗努恩LadnunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ladnun",
		TitleName: "罗努恩",
		TitleCode: "b_ladnun",
	}
}
