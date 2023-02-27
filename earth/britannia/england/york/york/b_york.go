package york

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约克YorkBarony struct {
	feud.BaseBarony
}

var BYork约克 feud.Barony = &约克YorkBarony{}

func init() {
    f := BYork约克.(*约克YorkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "york",
		TitleName: "约克",
		TitleCode: "b_york",
	}
}
