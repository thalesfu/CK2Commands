package baalbek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎赫勒ZahleBarony struct {
	feud.BaseBarony
}

var BZahle扎赫勒 feud.Barony = &扎赫勒ZahleBarony{}

func init() {
    f := BZahle扎赫勒.(*扎赫勒ZahleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zahle",
		TitleName: "扎赫勒",
		TitleCode: "b_zahle",
	}
}
