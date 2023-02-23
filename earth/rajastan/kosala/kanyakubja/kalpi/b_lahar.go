package kalpi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉诃罗LaharBarony struct {
	feud.BaseBarony
}

var BLahar拉诃罗 feud.Barony = &拉诃罗LaharBarony{}

func init() {
	f := BLahar拉诃罗.(*拉诃罗LaharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lahar",
		TitleName: "拉诃罗",
		TitleCode: "b_lahar",
	}
}
