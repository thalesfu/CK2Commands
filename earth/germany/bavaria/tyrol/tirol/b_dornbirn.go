package tirol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多恩比恩DornbirnBarony struct {
	feud.BaseBarony
}

var BDornbirn多恩比恩 feud.Barony = &多恩比恩DornbirnBarony{}

func init() {
    f := BDornbirn多恩比恩.(*多恩比恩DornbirnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dornbirn",
		TitleName: "多恩比恩",
		TitleCode: "b_dornbirn",
	}
}
