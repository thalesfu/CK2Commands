package bumthang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖恩ZugneBarony struct {
	feud.BaseBarony
}

var BZugne祖恩 feud.Barony = &祖恩ZugneBarony{}

func init() {
    f := BZugne祖恩.(*祖恩ZugneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zugne",
		TitleName: "祖恩",
		TitleCode: "b_zugne",
	}
}
