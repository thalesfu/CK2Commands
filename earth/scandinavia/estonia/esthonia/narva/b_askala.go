package narva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯凯莱AskalaBarony struct {
	feud.BaseBarony
}

var BAskala埃斯凯莱 feud.Barony = &埃斯凯莱AskalaBarony{}

func init() {
	f := BAskala埃斯凯莱.(*埃斯凯莱AskalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "askala",
		TitleName: "埃斯凯莱",
		TitleCode: "b_askala",
	}
}
