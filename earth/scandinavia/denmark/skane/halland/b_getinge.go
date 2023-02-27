package halland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶廷厄GetingeBarony struct {
	feud.BaseBarony
}

var BGetinge耶廷厄 feud.Barony = &耶廷厄GetingeBarony{}

func init() {
    f := BGetinge耶廷厄.(*耶廷厄GetingeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "getinge",
		TitleName: "耶廷厄",
		TitleCode: "b_getinge",
	}
}
