package madaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖斯泰勒QastalBarony struct {
	feud.BaseBarony
}

var BQastal盖斯泰勒 feud.Barony = &盖斯泰勒QastalBarony{}

func init() {
	f := BQastal盖斯泰勒.(*盖斯泰勒QastalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qastal",
		TitleName: "盖斯泰勒",
		TitleCode: "b_qastal",
	}
}
