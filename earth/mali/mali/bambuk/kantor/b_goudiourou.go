package kantor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古迪乌鲁GoudiourouBarony struct {
	feud.BaseBarony
}

var BGoudiourou古迪乌鲁 feud.Barony = &古迪乌鲁GoudiourouBarony{}

func init() {
    f := BGoudiourou古迪乌鲁.(*古迪乌鲁GoudiourouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goudiourou",
		TitleName: "古迪乌鲁",
		TitleCode: "b_goudiourou",
	}
}
