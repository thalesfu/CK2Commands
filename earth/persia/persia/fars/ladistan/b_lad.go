package ladistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉德LadBarony struct {
	feud.BaseBarony
}

var BLad拉德 feud.Barony = &拉德LadBarony{}

func init() {
	f := BLad拉德.(*拉德LadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lad",
		TitleName: "拉德",
		TitleCode: "b_lad",
	}
}
