package trapezous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里宗RiziniBarony struct {
	feud.BaseBarony
}

var BRizini里宗 feud.Barony = &里宗RiziniBarony{}

func init() {
    f := BRizini里宗.(*里宗RiziniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rizini",
		TitleName: "里宗",
		TitleCode: "b_rizini",
	}
}
