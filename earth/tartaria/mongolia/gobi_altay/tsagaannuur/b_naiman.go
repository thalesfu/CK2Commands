package tsagaannuur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乃蛮NaimanBarony struct {
	feud.BaseBarony
}

var BNaiman乃蛮 feud.Barony = &乃蛮NaimanBarony{}

func init() {
    f := BNaiman乃蛮.(*乃蛮NaimanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naiman",
		TitleName: "乃蛮",
		TitleCode: "b_naiman",
	}
}
