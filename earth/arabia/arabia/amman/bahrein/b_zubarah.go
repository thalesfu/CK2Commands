package bahrein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖巴拉ZubarahBarony struct {
	feud.BaseBarony
}

var BZubarah祖巴拉 feud.Barony = &祖巴拉ZubarahBarony{}

func init() {
	f := BZubarah祖巴拉.(*祖巴拉ZubarahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zubarah",
		TitleName: "祖巴拉",
		TitleCode: "b_zubarah",
	}
}
