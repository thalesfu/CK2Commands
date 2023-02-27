package novgorod_seversky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴赫马奇BakhmatchBarony struct {
	feud.BaseBarony
}

var BBakhmatch巴赫马奇 feud.Barony = &巴赫马奇BakhmatchBarony{}

func init() {
    f := BBakhmatch巴赫马奇.(*巴赫马奇BakhmatchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakhmatch",
		TitleName: "巴赫马奇",
		TitleCode: "b_bakhmatch",
	}
}
