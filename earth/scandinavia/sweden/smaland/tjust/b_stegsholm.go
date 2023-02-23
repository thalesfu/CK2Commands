package tjust

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特格斯霍尔姆StegsholmBarony struct {
	feud.BaseBarony
}

var BStegsholm斯特格斯霍尔姆 feud.Barony = &斯特格斯霍尔姆StegsholmBarony{}

func init() {
	f := BStegsholm斯特格斯霍尔姆.(*斯特格斯霍尔姆StegsholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stegsholm",
		TitleName: "斯特格斯霍尔姆",
		TitleCode: "b_stegsholm",
	}
}
