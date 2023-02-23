package venezia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕莱斯特里纳PallestrinaBarony struct {
	feud.BaseBarony
}

var BPallestrina帕莱斯特里纳 feud.Barony = &帕莱斯特里纳PallestrinaBarony{}

func init() {
	f := BPallestrina帕莱斯特里纳.(*帕莱斯特里纳PallestrinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pallestrina",
		TitleName: "帕莱斯特里纳",
		TitleCode: "b_pallestrina",
	}
}
