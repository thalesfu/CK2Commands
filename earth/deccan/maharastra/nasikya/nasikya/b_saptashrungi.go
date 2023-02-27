package nasikya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 七峰SaptashrungiBarony struct {
	feud.BaseBarony
}

var BSaptashrungi七峰 feud.Barony = &七峰SaptashrungiBarony{}

func init() {
    f := BSaptashrungi七峰.(*七峰SaptashrungiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saptashrungi",
		TitleName: "七峰",
		TitleCode: "b_saptashrungi",
	}
}
