package nordgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔施泰因ErsteinBarony struct {
	feud.BaseBarony
}

var BErstein埃尔施泰因 feud.Barony = &埃尔施泰因ErsteinBarony{}

func init() {
	f := BErstein埃尔施泰因.(*埃尔施泰因ErsteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erstein",
		TitleName: "埃尔施泰因",
		TitleCode: "b_erstein",
	}
}
