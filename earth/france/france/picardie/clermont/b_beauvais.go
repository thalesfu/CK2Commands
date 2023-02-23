package clermont

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博韦BeauvaisBarony struct {
	feud.BaseBarony
}

var BBeauvais博韦 feud.Barony = &博韦BeauvaisBarony{}

func init() {
	f := BBeauvais博韦.(*博韦BeauvaisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beauvais",
		TitleName: "博韦",
		TitleCode: "b_beauvais",
	}
}
