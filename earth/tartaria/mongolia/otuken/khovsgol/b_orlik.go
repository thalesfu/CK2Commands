package khovsgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔利克OrlikBarony struct {
	feud.BaseBarony
}

var BOrlik奥尔利克 feud.Barony = &奥尔利克OrlikBarony{}

func init() {
	f := BOrlik奥尔利克.(*奥尔利克OrlikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orlik",
		TitleName: "奥尔利克",
		TitleCode: "b_orlik",
	}
}
