package almaty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨雷耶西克SaryesikBarony struct {
	feud.BaseBarony
}

var BSaryesik萨雷耶西克 feud.Barony = &萨雷耶西克SaryesikBarony{}

func init() {
	f := BSaryesik萨雷耶西克.(*萨雷耶西克SaryesikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saryesik",
		TitleName: "萨雷耶西克",
		TitleCode: "b_saryesik",
	}
}
