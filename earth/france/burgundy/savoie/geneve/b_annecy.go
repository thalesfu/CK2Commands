package geneve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿讷西AnnecyBarony struct {
	feud.BaseBarony
}

var BAnnecy阿讷西 feud.Barony = &阿讷西AnnecyBarony{}

func init() {
	f := BAnnecy阿讷西.(*阿讷西AnnecyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "annecy",
		TitleName: "阿讷西",
		TitleCode: "b_annecy",
	}
}
