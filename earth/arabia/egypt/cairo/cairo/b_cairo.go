package cairo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 开罗CairoBarony struct {
	feud.BaseBarony
}

var BCairo开罗 feud.Barony = &开罗CairoBarony{}

func init() {
    f := BCairo开罗.(*开罗CairoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cairo",
		TitleName: "开罗",
		TitleCode: "b_cairo",
	}
}
