package kermanshah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕韦PavehBarony struct {
	feud.BaseBarony
}

var BPaveh帕韦 feud.Barony = &帕韦PavehBarony{}

func init() {
    f := BPaveh帕韦.(*帕韦PavehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paveh",
		TitleName: "帕韦",
		TitleCode: "b_paveh",
	}
}
