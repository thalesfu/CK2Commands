package tirol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰德克LandeckBarony struct {
	feud.BaseBarony
}

var BLandeck兰德克 feud.Barony = &兰德克LandeckBarony{}

func init() {
	f := BLandeck兰德克.(*兰德克LandeckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "landeck",
		TitleName: "兰德克",
		TitleCode: "b_landeck",
	}
}
