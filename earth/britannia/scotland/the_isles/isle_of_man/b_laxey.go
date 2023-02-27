package isle_of_man

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉克西LaxeyBarony struct {
	feud.BaseBarony
}

var BLaxey拉克西 feud.Barony = &拉克西LaxeyBarony{}

func init() {
    f := BLaxey拉克西.(*拉克西LaxeyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laxey",
		TitleName: "拉克西",
		TitleCode: "b_laxey",
	}
}
