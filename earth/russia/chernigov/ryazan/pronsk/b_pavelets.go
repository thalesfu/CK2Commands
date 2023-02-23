package pronsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕韦列茨PaveletsBarony struct {
	feud.BaseBarony
}

var BPavelets帕韦列茨 feud.Barony = &帕韦列茨PaveletsBarony{}

func init() {
	f := BPavelets帕韦列茨.(*帕韦列茨PaveletsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pavelets",
		TitleName: "帕韦列茨",
		TitleCode: "b_pavelets",
	}
}
