package abosyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰讷斯RandersBarony struct {
	feud.BaseBarony
}

var BRanders兰讷斯 feud.Barony = &兰讷斯RandersBarony{}

func init() {
    f := BRanders兰讷斯.(*兰讷斯RandersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "randers",
		TitleName: "兰讷斯",
		TitleCode: "b_randers",
	}
}
