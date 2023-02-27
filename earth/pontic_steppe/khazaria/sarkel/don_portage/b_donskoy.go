package don_portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 顿斯科伊DonskoyBarony struct {
	feud.BaseBarony
}

var BDonskoy顿斯科伊 feud.Barony = &顿斯科伊DonskoyBarony{}

func init() {
    f := BDonskoy顿斯科伊.(*顿斯科伊DonskoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "donskoy",
		TitleName: "顿斯科伊",
		TitleCode: "b_donskoy",
	}
}
