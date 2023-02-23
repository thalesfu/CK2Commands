package bayda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿塔格AtaqBarony struct {
	feud.BaseBarony
}

var BAtaq阿塔格 feud.Barony = &阿塔格AtaqBarony{}

func init() {
	f := BAtaq阿塔格.(*阿塔格AtaqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ataq",
		TitleName: "阿塔格",
		TitleCode: "b_ataq",
	}
}
