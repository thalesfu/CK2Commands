package abkhazia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖格迪迪ZugdidiBarony struct {
	feud.BaseBarony
}

var BZugdidi祖格迪迪 feud.Barony = &祖格迪迪ZugdidiBarony{}

func init() {
	f := BZugdidi祖格迪迪.(*祖格迪迪ZugdidiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zugdidi",
		TitleName: "祖格迪迪",
		TitleCode: "b_zugdidi",
	}
}
