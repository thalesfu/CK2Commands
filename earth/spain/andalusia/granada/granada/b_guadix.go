package granada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓜迪克斯GuadixBarony struct {
	feud.BaseBarony
}

var BGuadix瓜迪克斯 feud.Barony = &瓜迪克斯GuadixBarony{}

func init() {
	f := BGuadix瓜迪克斯.(*瓜迪克斯GuadixBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guadix",
		TitleName: "瓜迪克斯",
		TitleCode: "b_guadix",
	}
}
