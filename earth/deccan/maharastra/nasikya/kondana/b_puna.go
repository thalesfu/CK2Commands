package kondana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普纳PunaBarony struct {
	feud.BaseBarony
}

var BPuna普纳 feud.Barony = &普纳PunaBarony{}

func init() {
    f := BPuna普纳.(*普纳PunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "puna",
		TitleName: "普纳",
		TitleCode: "b_puna",
	}
}
