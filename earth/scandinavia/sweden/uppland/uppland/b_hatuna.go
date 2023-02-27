package uppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍图纳HatunaBarony struct {
	feud.BaseBarony
}

var BHatuna霍图纳 feud.Barony = &霍图纳HatunaBarony{}

func init() {
    f := BHatuna霍图纳.(*霍图纳HatunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hatuna",
		TitleName: "霍图纳",
		TitleCode: "b_hatuna",
	}
}
