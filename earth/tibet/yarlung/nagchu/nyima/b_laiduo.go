package nyima

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 来多LaiduoBarony struct {
	feud.BaseBarony
}

var BLaiduo来多 feud.Barony = &来多LaiduoBarony{}

func init() {
	f := BLaiduo来多.(*来多LaiduoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laiduo",
		TitleName: "来多",
		TitleCode: "b_laiduo",
	}
}
