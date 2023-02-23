package ulster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉恩LarneBarony struct {
	feud.BaseBarony
}

var BLarne拉恩 feud.Barony = &拉恩LarneBarony{}

func init() {
	f := BLarne拉恩.(*拉恩LarneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "larne",
		TitleName: "拉恩",
		TitleCode: "b_larne",
	}
}
