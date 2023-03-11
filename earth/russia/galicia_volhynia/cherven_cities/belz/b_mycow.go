package belz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅楚夫MycowBarony struct {
	feud.BaseBarony
}

var BMycow梅楚夫 feud.Barony = &梅楚夫MycowBarony{}

func init() {
    f := BMycow梅楚夫.(*梅楚夫MycowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mycow",
		TitleName: "梅楚夫",
		TitleCode: "b_mycow",
	}
}
