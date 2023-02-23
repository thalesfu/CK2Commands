package kuwait

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富海希勒AlfalalheelBarony struct {
	feud.BaseBarony
}

var BAlfalalheel富海希勒 feud.Barony = &富海希勒AlfalalheelBarony{}

func init() {
	f := BAlfalalheel富海希勒.(*富海希勒AlfalalheelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alfalalheel",
		TitleName: "富海希勒",
		TitleCode: "b_alfalalheel",
	}
}
