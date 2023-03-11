package zvenyhorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布茨克ButskBarony struct {
	feud.BaseBarony
}

var BButsk布茨克 feud.Barony = &布茨克ButskBarony{}

func init() {
    f := BButsk布茨克.(*布茨克ButskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "butsk",
		TitleName: "布茨克",
		TitleCode: "b_butsk",
	}
}
