package rosello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞雷特CeretBarony struct {
	feud.BaseBarony
}

var BCeret塞雷特 feud.Barony = &塞雷特CeretBarony{}

func init() {
	f := BCeret塞雷特.(*塞雷特CeretBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ceret",
		TitleName: "塞雷特",
		TitleCode: "b_ceret",
	}
}
