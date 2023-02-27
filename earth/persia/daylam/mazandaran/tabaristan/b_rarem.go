package tabaristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉雷姆RaremBarony struct {
	feud.BaseBarony
}

var BRarem拉雷姆 feud.Barony = &拉雷姆RaremBarony{}

func init() {
    f := BRarem拉雷姆.(*拉雷姆RaremBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rarem",
		TitleName: "拉雷姆",
		TitleCode: "b_rarem",
	}
}
