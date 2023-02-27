package slutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列奇察RechytsaBarony struct {
	feud.BaseBarony
}

var BRechytsa列奇察 feud.Barony = &列奇察RechytsaBarony{}

func init() {
    f := BRechytsa列奇察.(*列奇察RechytsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rechytsa",
		TitleName: "列奇察",
		TitleCode: "b_rechytsa",
	}
}
