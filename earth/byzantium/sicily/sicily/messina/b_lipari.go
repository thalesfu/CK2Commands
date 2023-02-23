package messina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利帕里LipariBarony struct {
	feud.BaseBarony
}

var BLipari利帕里 feud.Barony = &利帕里LipariBarony{}

func init() {
	f := BLipari利帕里.(*利帕里LipariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lipari",
		TitleName: "利帕里",
		TitleCode: "b_lipari",
	}
}
