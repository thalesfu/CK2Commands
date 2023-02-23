package foggia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福贾FoggiaBarony struct {
	feud.BaseBarony
}

var BFoggia福贾 feud.Barony = &福贾FoggiaBarony{}

func init() {
	f := BFoggia福贾.(*福贾FoggiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "foggia",
		TitleName: "福贾",
		TitleCode: "b_foggia",
	}
}
