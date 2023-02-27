package bandhugadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆呼利般陀BahoribandBarony struct {
	feud.BaseBarony
}

var BBahoriband婆呼利般陀 feud.Barony = &婆呼利般陀BahoribandBarony{}

func init() {
    f := BBahoriband婆呼利般陀.(*婆呼利般陀BahoribandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bahoriband",
		TitleName: "婆呼利般陀",
		TitleCode: "b_bahoriband",
	}
}
