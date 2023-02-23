package thrake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥莱梯古斯AulaeitichusBarony struct {
	feud.BaseBarony
}

var BAulaeitichus奥莱梯古斯 feud.Barony = &奥莱梯古斯AulaeitichusBarony{}

func init() {
	f := BAulaeitichus奥莱梯古斯.(*奥莱梯古斯AulaeitichusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aulaeitichus",
		TitleName: "奥莱梯古斯",
		TitleCode: "b_aulaeitichus",
	}
}
