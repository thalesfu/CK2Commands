package thisageta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维扎伊VizhayBarony struct {
	feud.BaseBarony
}

var BVizhay维扎伊 feud.Barony = &维扎伊VizhayBarony{}

func init() {
    f := BVizhay维扎伊.(*维扎伊VizhayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vizhay",
		TitleName: "维扎伊",
		TitleCode: "b_vizhay",
	}
}
