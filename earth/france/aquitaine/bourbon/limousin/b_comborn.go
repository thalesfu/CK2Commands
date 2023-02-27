package limousin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔博恩CombornBarony struct {
	feud.BaseBarony
}

var BComborn孔博恩 feud.Barony = &孔博恩CombornBarony{}

func init() {
    f := BComborn孔博恩.(*孔博恩CombornBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "comborn",
		TitleName: "孔博恩",
		TitleCode: "b_comborn",
	}
}
