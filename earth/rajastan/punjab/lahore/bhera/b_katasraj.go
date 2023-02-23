package bhera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泪眼KatasrajBarony struct {
	feud.BaseBarony
}

var BKatasraj泪眼 feud.Barony = &泪眼KatasrajBarony{}

func init() {
	f := BKatasraj泪眼.(*泪眼KatasrajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katasraj",
		TitleName: "泪眼",
		TitleCode: "b_katasraj",
	}
}
