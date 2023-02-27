package albret

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米米藏MimizanBarony struct {
	feud.BaseBarony
}

var BMimizan米米藏 feud.Barony = &米米藏MimizanBarony{}

func init() {
    f := BMimizan米米藏.(*米米藏MimizanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mimizan",
		TitleName: "米米藏",
		TitleCode: "b_mimizan",
	}
}
