package marienburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马林韦尔德MarienwerderBarony struct {
	feud.BaseBarony
}

var BMarienwerder马林韦尔德 feud.Barony = &马林韦尔德MarienwerderBarony{}

func init() {
    f := BMarienwerder马林韦尔德.(*马林韦尔德MarienwerderBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marienwerder",
		TitleName: "马林韦尔德",
		TitleCode: "b_marienwerder",
	}
}
