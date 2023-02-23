package bayda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝达BaydaBarony struct {
	feud.BaseBarony
}

var BBayda贝达 feud.Barony = &贝达BaydaBarony{}

func init() {
	f := BBayda贝达.(*贝达BaydaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bayda",
		TitleName: "贝达",
		TitleCode: "b_bayda",
	}
}
