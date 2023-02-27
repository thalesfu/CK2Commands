package slupsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯武普斯克SlupskBarony struct {
	feud.BaseBarony
}

var BSlupsk斯武普斯克 feud.Barony = &斯武普斯克SlupskBarony{}

func init() {
    f := BSlupsk斯武普斯克.(*斯武普斯克SlupskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slupsk",
		TitleName: "斯武普斯克",
		TitleCode: "b_slupsk",
	}
}
