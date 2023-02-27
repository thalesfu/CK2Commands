package lower_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希维博津SwiebozinBarony struct {
	feud.BaseBarony
}

var BSwiebozin希维博津 feud.Barony = &希维博津SwiebozinBarony{}

func init() {
    f := BSwiebozin希维博津.(*希维博津SwiebozinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "swiebozin",
		TitleName: "希维博津",
		TitleCode: "b_swiebozin",
	}
}
