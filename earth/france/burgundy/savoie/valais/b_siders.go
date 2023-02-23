package valais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西德尔斯SidersBarony struct {
	feud.BaseBarony
}

var BSiders西德尔斯 feud.Barony = &西德尔斯SidersBarony{}

func init() {
	f := BSiders西德尔斯.(*西德尔斯SidersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siders",
		TitleName: "西德尔斯",
		TitleCode: "b_siders",
	}
}
