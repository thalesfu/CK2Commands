package moray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈恩NairnBarony struct {
	feud.BaseBarony
}

var BNairn奈恩 feud.Barony = &奈恩NairnBarony{}

func init() {
	f := BNairn奈恩.(*奈恩NairnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nairn",
		TitleName: "奈恩",
		TitleCode: "b_nairn",
	}
}
