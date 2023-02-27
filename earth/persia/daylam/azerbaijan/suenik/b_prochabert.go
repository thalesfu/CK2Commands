package suenik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉洛斯比尔特ProchabertBarony struct {
	feud.BaseBarony
}

var BProchabert拉洛斯比尔特 feud.Barony = &拉洛斯比尔特ProchabertBarony{}

func init() {
    f := BProchabert拉洛斯比尔特.(*拉洛斯比尔特ProchabertBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prochabert",
		TitleName: "拉洛斯比尔特",
		TitleCode: "b_prochabert",
	}
}
