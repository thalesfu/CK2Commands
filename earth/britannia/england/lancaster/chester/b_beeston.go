package chester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比斯顿BeestonBarony struct {
	feud.BaseBarony
}

var BBeeston比斯顿 feud.Barony = &比斯顿BeestonBarony{}

func init() {
	f := BBeeston比斯顿.(*比斯顿BeestonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beeston",
		TitleName: "比斯顿",
		TitleCode: "b_beeston",
	}
}
