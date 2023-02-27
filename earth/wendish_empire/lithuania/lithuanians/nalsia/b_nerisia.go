package nalsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内里西亚NerisiaBarony struct {
	feud.BaseBarony
}

var BNerisia内里西亚 feud.Barony = &内里西亚NerisiaBarony{}

func init() {
    f := BNerisia内里西亚.(*内里西亚NerisiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nerisia",
		TitleName: "内里西亚",
		TitleCode: "b_nerisia",
	}
}
