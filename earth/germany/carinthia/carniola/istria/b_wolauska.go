package istria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃劳斯卡WolauskaBarony struct {
	feud.BaseBarony
}

var BWolauska沃劳斯卡 feud.Barony = &沃劳斯卡WolauskaBarony{}

func init() {
    f := BWolauska沃劳斯卡.(*沃劳斯卡WolauskaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wolauska",
		TitleName: "沃劳斯卡",
		TitleCode: "b_wolauska",
	}
}
