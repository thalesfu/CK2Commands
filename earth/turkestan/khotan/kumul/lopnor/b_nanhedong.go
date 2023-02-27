package lopnor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 南河东NanhedongBarony struct {
	feud.BaseBarony
}

var BNanhedong南河东 feud.Barony = &南河东NanhedongBarony{}

func init() {
    f := BNanhedong南河东.(*南河东NanhedongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nanhedong",
		TitleName: "南河东",
		TitleCode: "b_nanhedong",
	}
}
