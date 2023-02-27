package bedford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿什里奇AshridgeBarony struct {
	feud.BaseBarony
}

var BAshridge阿什里奇 feud.Barony = &阿什里奇AshridgeBarony{}

func init() {
    f := BAshridge阿什里奇.(*阿什里奇AshridgeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ashridge",
		TitleName: "阿什里奇",
		TitleCode: "b_ashridge",
	}
}
