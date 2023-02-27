package istakhr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯塔赫尔IstakhrBarony struct {
	feud.BaseBarony
}

var BIstakhr伊斯塔赫尔 feud.Barony = &伊斯塔赫尔IstakhrBarony{}

func init() {
    f := BIstakhr伊斯塔赫尔.(*伊斯塔赫尔IstakhrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "istakhr",
		TitleName: "伊斯塔赫尔",
		TitleCode: "b_istakhr",
	}
}
