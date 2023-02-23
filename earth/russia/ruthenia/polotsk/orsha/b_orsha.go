package orsha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔沙OrshaBarony struct {
	feud.BaseBarony
}

var BOrsha奥尔沙 feud.Barony = &奥尔沙OrshaBarony{}

func init() {
	f := BOrsha奥尔沙.(*奥尔沙OrshaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orsha",
		TitleName: "奥尔沙",
		TitleCode: "b_orsha",
	}
}
