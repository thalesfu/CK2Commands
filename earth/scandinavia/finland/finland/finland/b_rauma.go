package finland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳马RaumaBarony struct {
	feud.BaseBarony
}

var BRauma劳马 feud.Barony = &劳马RaumaBarony{}

func init() {
	f := BRauma劳马.(*劳马RaumaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rauma",
		TitleName: "劳马",
		TitleCode: "b_rauma",
	}
}
