package byzantion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 潘普顿PemptonBarony struct {
	feud.BaseBarony
}

var BPempton潘普顿 feud.Barony = &潘普顿PemptonBarony{}

func init() {
	f := BPempton潘普顿.(*潘普顿PemptonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pempton",
		TitleName: "潘普顿",
		TitleCode: "b_pempton",
	}
}
