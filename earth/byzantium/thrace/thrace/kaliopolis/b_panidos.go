package kaliopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕尼多斯PanidosBarony struct {
	feud.BaseBarony
}

var BPanidos帕尼多斯 feud.Barony = &帕尼多斯PanidosBarony{}

func init() {
	f := BPanidos帕尼多斯.(*帕尼多斯PanidosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "panidos",
		TitleName: "帕尼多斯",
		TitleCode: "b_panidos",
	}
}
