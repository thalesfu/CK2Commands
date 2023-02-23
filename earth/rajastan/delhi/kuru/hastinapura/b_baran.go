package hastinapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆兰BaranBarony struct {
	feud.BaseBarony
}

var BBaran婆兰 feud.Barony = &婆兰BaranBarony{}

func init() {
	f := BBaran婆兰.(*婆兰BaranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baran",
		TitleName: "婆兰",
		TitleCode: "b_baran",
	}
}
