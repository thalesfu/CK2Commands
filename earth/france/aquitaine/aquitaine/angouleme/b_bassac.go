package angouleme

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴萨克BassacBarony struct {
	feud.BaseBarony
}

var BBassac巴萨克 feud.Barony = &巴萨克BassacBarony{}

func init() {
	f := BBassac巴萨克.(*巴萨克BassacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bassac",
		TitleName: "巴萨克",
		TitleCode: "b_bassac",
	}
}
