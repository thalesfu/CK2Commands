package shiraz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶斯塔赫尔EstakhrBarony struct {
	feud.BaseBarony
}

var BEstakhr耶斯塔赫尔 feud.Barony = &耶斯塔赫尔EstakhrBarony{}

func init() {
    f := BEstakhr耶斯塔赫尔.(*耶斯塔赫尔EstakhrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "estakhr",
		TitleName: "耶斯塔赫尔",
		TitleCode: "b_estakhr",
	}
}
