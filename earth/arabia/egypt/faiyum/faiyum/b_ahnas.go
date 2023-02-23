package faiyum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿赫纳斯AhnasBarony struct {
	feud.BaseBarony
}

var BAhnas阿赫纳斯 feud.Barony = &阿赫纳斯AhnasBarony{}

func init() {
	f := BAhnas阿赫纳斯.(*阿赫纳斯AhnasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ahnas",
		TitleName: "阿赫纳斯",
		TitleCode: "b_ahnas",
	}
}
