package trapezous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里泽翁RizaionBarony struct {
	feud.BaseBarony
}

var BRizaion里泽翁 feud.Barony = &里泽翁RizaionBarony{}

func init() {
	f := BRizaion里泽翁.(*里泽翁RizaionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rizaion",
		TitleName: "里泽翁",
		TitleCode: "b_rizaion",
	}
}
