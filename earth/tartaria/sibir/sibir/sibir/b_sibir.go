package sibir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 失必儿SibirBarony struct {
	feud.BaseBarony
}

var BSibir失必儿 feud.Barony = &失必儿SibirBarony{}

func init() {
    f := BSibir失必儿.(*失必儿SibirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sibir",
		TitleName: "失必儿",
		TitleCode: "b_sibir",
	}
}
