package nordland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒德于RodoyBarony struct {
	feud.BaseBarony
}

var BRodoy勒德于 feud.Barony = &勒德于RodoyBarony{}

func init() {
    f := BRodoy勒德于.(*勒德于RodoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rodoy",
		TitleName: "勒德于",
		TitleCode: "b_rodoy",
	}
}
