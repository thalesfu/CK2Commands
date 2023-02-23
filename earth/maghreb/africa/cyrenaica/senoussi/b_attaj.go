package senoussi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔季AttajBarony struct {
	feud.BaseBarony
}

var BAttaj塔季 feud.Barony = &塔季AttajBarony{}

func init() {
	f := BAttaj塔季.(*塔季AttajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "attaj",
		TitleName: "塔季",
		TitleCode: "b_attaj",
	}
}
