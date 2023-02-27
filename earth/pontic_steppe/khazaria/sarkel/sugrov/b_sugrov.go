package sugrov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏格罗夫SugrovBarony struct {
	feud.BaseBarony
}

var BSugrov苏格罗夫 feud.Barony = &苏格罗夫SugrovBarony{}

func init() {
    f := BSugrov苏格罗夫.(*苏格罗夫SugrovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sugrov",
		TitleName: "苏格罗夫",
		TitleCode: "b_sugrov",
	}
}
