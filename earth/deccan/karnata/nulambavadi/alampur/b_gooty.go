package alampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古蒂GootyBarony struct {
	feud.BaseBarony
}

var BGooty古蒂 feud.Barony = &古蒂GootyBarony{}

func init() {
    f := BGooty古蒂.(*古蒂GootyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gooty",
		TitleName: "古蒂",
		TitleCode: "b_gooty",
	}
}
