package dunhuang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 方盘FangpanBarony struct {
	feud.BaseBarony
}

var BFangpan方盘 feud.Barony = &方盘FangpanBarony{}

func init() {
	f := BFangpan方盘.(*方盘FangpanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fangpan",
		TitleName: "方盘",
		TitleCode: "b_fangpan",
	}
}
