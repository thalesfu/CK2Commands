package dijon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥坦AutunBarony struct {
	feud.BaseBarony
}

var BAutun奥坦 feud.Barony = &奥坦AutunBarony{}

func init() {
    f := BAutun奥坦.(*奥坦AutunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "autun",
		TitleName: "奥坦",
		TitleCode: "b_autun",
	}
}
