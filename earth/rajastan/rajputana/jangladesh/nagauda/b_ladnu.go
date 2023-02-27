package nagauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗努LadnuBarony struct {
	feud.BaseBarony
}

var BLadnu罗努 feud.Barony = &罗努LadnuBarony{}

func init() {
    f := BLadnu罗努.(*罗努LadnuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ladnu",
		TitleName: "罗努",
		TitleCode: "b_ladnu",
	}
}
