package nilagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 潘加洛PanagalluBarony struct {
	feud.BaseBarony
}

var BPanagallu潘加洛 feud.Barony = &潘加洛PanagalluBarony{}

func init() {
	f := BPanagallu潘加洛.(*潘加洛PanagalluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "panagallu",
		TitleName: "潘加洛",
		TitleCode: "b_panagallu",
	}
}
