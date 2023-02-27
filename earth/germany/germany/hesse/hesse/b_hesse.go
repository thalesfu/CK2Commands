package hesse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑森HesseBarony struct {
	feud.BaseBarony
}

var BHesse黑森 feud.Barony = &黑森HesseBarony{}

func init() {
    f := BHesse黑森.(*黑森HesseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hesse",
		TitleName: "黑森",
		TitleCode: "b_hesse",
	}
}
