package alexandria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚历山大AlexandriaBarony struct {
	feud.BaseBarony
}

var BAlexandria亚历山大 feud.Barony = &亚历山大AlexandriaBarony{}

func init() {
	f := BAlexandria亚历山大.(*亚历山大AlexandriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alexandria",
		TitleName: "亚历山大",
		TitleCode: "b_alexandria",
	}
}
