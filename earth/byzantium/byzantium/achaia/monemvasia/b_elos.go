package monemvasia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾洛斯ElosBarony struct {
	feud.BaseBarony
}

var BElos艾洛斯 feud.Barony = &艾洛斯ElosBarony{}

func init() {
	f := BElos艾洛斯.(*艾洛斯ElosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elos",
		TitleName: "艾洛斯",
		TitleCode: "b_elos",
	}
}
