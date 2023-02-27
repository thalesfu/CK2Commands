package halland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉霍尔姆LaholmBarony struct {
	feud.BaseBarony
}

var BLaholm拉霍尔姆 feud.Barony = &拉霍尔姆LaholmBarony{}

func init() {
    f := BLaholm拉霍尔姆.(*拉霍尔姆LaholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laholm",
		TitleName: "拉霍尔姆",
		TitleCode: "b_laholm",
	}
}
