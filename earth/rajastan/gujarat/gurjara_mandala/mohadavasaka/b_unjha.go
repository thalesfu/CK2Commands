package mohadavasaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 温恰UnjhaBarony struct {
	feud.BaseBarony
}

var BUnjha温恰 feud.Barony = &温恰UnjhaBarony{}

func init() {
    f := BUnjha温恰.(*温恰UnjhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "unjha",
		TitleName: "温恰",
		TitleCode: "b_unjha",
	}
}
