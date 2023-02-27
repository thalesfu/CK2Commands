package zamindawar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰吉LangiBarony struct {
	feud.BaseBarony
}

var BLangi兰吉 feud.Barony = &兰吉LangiBarony{}

func init() {
    f := BLangi兰吉.(*兰吉LangiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "langi",
		TitleName: "兰吉",
		TitleCode: "b_langi",
	}
}
