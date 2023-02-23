package muztau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄斯克门OskemenBarony struct {
	feud.BaseBarony
}

var BOskemen厄斯克门 feud.Barony = &厄斯克门OskemenBarony{}

func init() {
	f := BOskemen厄斯克门.(*厄斯克门OskemenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oskemen",
		TitleName: "厄斯克门",
		TitleCode: "b_oskemen",
	}
}
