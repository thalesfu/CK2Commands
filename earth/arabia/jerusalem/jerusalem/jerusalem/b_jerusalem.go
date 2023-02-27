package jerusalem

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶路撒冷JerusalemBarony struct {
	feud.BaseBarony
}

var BJerusalem耶路撒冷 feud.Barony = &耶路撒冷JerusalemBarony{}

func init() {
    f := BJerusalem耶路撒冷.(*耶路撒冷JerusalemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jerusalem",
		TitleName: "耶路撒冷",
		TitleCode: "b_jerusalem",
	}
}
