package nikopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯库斯OescusBarony struct {
	feud.BaseBarony
}

var BOescus埃斯库斯 feud.Barony = &埃斯库斯OescusBarony{}

func init() {
    f := BOescus埃斯库斯.(*埃斯库斯OescusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oescus",
		TitleName: "埃斯库斯",
		TitleCode: "b_oescus",
	}
}
