package baghdad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉蒂菲亚LatifiyaBarony struct {
	feud.BaseBarony
}

var BLatifiya拉蒂菲亚 feud.Barony = &拉蒂菲亚LatifiyaBarony{}

func init() {
	f := BLatifiya拉蒂菲亚.(*拉蒂菲亚LatifiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "latifiya",
		TitleName: "拉蒂菲亚",
		TitleCode: "b_latifiya",
	}
}
