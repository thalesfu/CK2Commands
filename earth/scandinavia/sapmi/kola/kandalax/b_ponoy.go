package kandalax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波诺伊PonoyBarony struct {
	feud.BaseBarony
}

var BPonoy波诺伊 feud.Barony = &波诺伊PonoyBarony{}

func init() {
    f := BPonoy波诺伊.(*波诺伊PonoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ponoy",
		TitleName: "波诺伊",
		TitleCode: "b_ponoy",
	}
}
