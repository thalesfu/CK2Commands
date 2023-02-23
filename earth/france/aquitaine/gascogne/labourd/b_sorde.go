package labourd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索尔德SordeBarony struct {
	feud.BaseBarony
}

var BSorde索尔德 feud.Barony = &索尔德SordeBarony{}

func init() {
	f := BSorde索尔德.(*索尔德SordeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sorde",
		TitleName: "索尔德",
		TitleCode: "b_sorde",
	}
}
