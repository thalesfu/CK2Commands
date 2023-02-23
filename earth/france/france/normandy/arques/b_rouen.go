package arques

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁昂RouenBarony struct {
	feud.BaseBarony
}

var BRouen鲁昂 feud.Barony = &鲁昂RouenBarony{}

func init() {
	f := BRouen鲁昂.(*鲁昂RouenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rouen",
		TitleName: "鲁昂",
		TitleCode: "b_rouen",
	}
}
