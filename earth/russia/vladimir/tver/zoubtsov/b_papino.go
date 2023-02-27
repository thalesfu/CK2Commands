package zoubtsov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕皮诺PapinoBarony struct {
	feud.BaseBarony
}

var BPapino帕皮诺 feud.Barony = &帕皮诺PapinoBarony{}

func init() {
    f := BPapino帕皮诺.(*帕皮诺PapinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "papino",
		TitleName: "帕皮诺",
		TitleCode: "b_papino",
	}
}
