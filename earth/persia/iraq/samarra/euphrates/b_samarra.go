package euphrates

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨迈拉SamarraBarony struct {
	feud.BaseBarony
}

var BSamarra萨迈拉 feud.Barony = &萨迈拉SamarraBarony{}

func init() {
	f := BSamarra萨迈拉.(*萨迈拉SamarraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samarra",
		TitleName: "萨迈拉",
		TitleCode: "b_samarra",
	}
}
