package negev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伏达特AvdatBarony struct {
	feud.BaseBarony
}

var BAvdat阿伏达特 feud.Barony = &阿伏达特AvdatBarony{}

func init() {
	f := BAvdat阿伏达特.(*阿伏达特AvdatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avdat",
		TitleName: "阿伏达特",
		TitleCode: "b_avdat",
	}
}
