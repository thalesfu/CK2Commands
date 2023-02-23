package samtho

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 若拉岗RolagangBarony struct {
	feud.BaseBarony
}

var BRolagang若拉岗 feud.Barony = &若拉岗RolagangBarony{}

func init() {
	f := BRolagang若拉岗.(*若拉岗RolagangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rolagang",
		TitleName: "若拉岗",
		TitleCode: "b_rolagang",
	}
}
