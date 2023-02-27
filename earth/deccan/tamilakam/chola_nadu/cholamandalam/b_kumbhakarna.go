package cholamandalam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宫婆羯罗拏KumbhakarnaBarony struct {
	feud.BaseBarony
}

var BKumbhakarna宫婆羯罗拏 feud.Barony = &宫婆羯罗拏KumbhakarnaBarony{}

func init() {
    f := BKumbhakarna宫婆羯罗拏.(*宫婆羯罗拏KumbhakarnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kumbhakarna",
		TitleName: "宫婆羯罗拏",
		TitleCode: "b_kumbhakarna",
	}
}
