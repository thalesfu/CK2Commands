package forez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣埃蒂安StetienneBarony struct {
	feud.BaseBarony
}

var BStetienne圣埃蒂安 feud.Barony = &圣埃蒂安StetienneBarony{}

func init() {
	f := BStetienne圣埃蒂安.(*圣埃蒂安StetienneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stetienne",
		TitleName: "圣埃蒂安",
		TitleCode: "b_stetienne",
	}
}
