package durdzukia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔东ArdonBarony struct {
	feud.BaseBarony
}

var BArdon阿尔东 feud.Barony = &阿尔东ArdonBarony{}

func init() {
    f := BArdon阿尔东.(*阿尔东ArdonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ardon",
		TitleName: "阿尔东",
		TitleCode: "b_ardon",
	}
}
