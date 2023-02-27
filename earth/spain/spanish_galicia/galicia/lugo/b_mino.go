package lugo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米尼奥MinoBarony struct {
	feud.BaseBarony
}

var BMino米尼奥 feud.Barony = &米尼奥MinoBarony{}

func init() {
    f := BMino米尼奥.(*米尼奥MinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mino",
		TitleName: "米尼奥",
		TitleCode: "b_mino",
	}
}
