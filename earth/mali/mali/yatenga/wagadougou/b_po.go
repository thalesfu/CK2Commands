package wagadougou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波PoBarony struct {
	feud.BaseBarony
}

var BPo波 feud.Barony = &波PoBarony{}

func init() {
    f := BPo波.(*波PoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "po",
		TitleName: "波",
		TitleCode: "b_po",
	}
}
