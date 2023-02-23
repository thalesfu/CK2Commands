package paro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洪热HungrelBarony struct {
	feud.BaseBarony
}

var BHungrel洪热 feud.Barony = &洪热HungrelBarony{}

func init() {
	f := BHungrel洪热.(*洪热HungrelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hungrel",
		TitleName: "洪热",
		TitleCode: "b_hungrel",
	}
}
