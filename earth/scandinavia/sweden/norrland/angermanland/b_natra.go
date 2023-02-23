package angermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅特拉NatraBarony struct {
	feud.BaseBarony
}

var BNatra涅特拉 feud.Barony = &涅特拉NatraBarony{}

func init() {
	f := BNatra涅特拉.(*涅特拉NatraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "natra",
		TitleName: "涅特拉",
		TitleCode: "b_natra",
	}
}
