package udabhanda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 濮达CharsadaBarony struct {
	feud.BaseBarony
}

var BCharsada濮达 feud.Barony = &濮达CharsadaBarony{}

func init() {
    f := BCharsada濮达.(*濮达CharsadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charsada",
		TitleName: "濮达",
		TitleCode: "b_charsada",
	}
}
