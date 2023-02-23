package bumthang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布姆唐BumthangBarony struct {
	feud.BaseBarony
}

var BBumthang布姆唐 feud.Barony = &布姆唐BumthangBarony{}

func init() {
	f := BBumthang布姆唐.(*布姆唐BumthangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bumthang",
		TitleName: "布姆唐",
		TitleCode: "b_bumthang",
	}
}
