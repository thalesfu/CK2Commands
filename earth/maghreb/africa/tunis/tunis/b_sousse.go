package tunis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏塞SousseBarony struct {
	feud.BaseBarony
}

var BSousse苏塞 feud.Barony = &苏塞SousseBarony{}

func init() {
	f := BSousse苏塞.(*苏塞SousseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sousse",
		TitleName: "苏塞",
		TitleCode: "b_sousse",
	}
}
