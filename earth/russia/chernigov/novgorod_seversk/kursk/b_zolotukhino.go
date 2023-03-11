package kursk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佐洛图希诺ZolotukhinoBarony struct {
	feud.BaseBarony
}

var BZolotukhino佐洛图希诺 feud.Barony = &佐洛图希诺ZolotukhinoBarony{}

func init() {
    f := BZolotukhino佐洛图希诺.(*佐洛图希诺ZolotukhinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zolotukhino",
		TitleName: "佐洛图希诺",
		TitleCode: "b_zolotukhino",
	}
}
