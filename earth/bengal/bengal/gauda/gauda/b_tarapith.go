package gauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多罗比他TarapithBarony struct {
	feud.BaseBarony
}

var BTarapith多罗比他 feud.Barony = &多罗比他TarapithBarony{}

func init() {
	f := BTarapith多罗比他.(*多罗比他TarapithBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarapith",
		TitleName: "多罗比他",
		TitleCode: "b_tarapith",
	}
}
