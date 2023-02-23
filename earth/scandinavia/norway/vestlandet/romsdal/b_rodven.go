package romsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒德文RodvenBarony struct {
	feud.BaseBarony
}

var BRodven勒德文 feud.Barony = &勒德文RodvenBarony{}

func init() {
	f := BRodven勒德文.(*勒德文RodvenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rodven",
		TitleName: "勒德文",
		TitleCode: "b_rodven",
	}
}
