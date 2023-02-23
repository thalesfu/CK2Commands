package baygal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔汗DarkhanBarony struct {
	feud.BaseBarony
}

var BDarkhan达尔汗 feud.Barony = &达尔汗DarkhanBarony{}

func init() {
	f := BDarkhan达尔汗.(*达尔汗DarkhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darkhan",
		TitleName: "达尔汗",
		TitleCode: "b_darkhan",
	}
}
