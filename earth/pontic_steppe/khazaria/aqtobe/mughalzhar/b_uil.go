package mughalzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌伊尔UilBarony struct {
	feud.BaseBarony
}

var BUil乌伊尔 feud.Barony = &乌伊尔UilBarony{}

func init() {
    f := BUil乌伊尔.(*乌伊尔UilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uil",
		TitleName: "乌伊尔",
		TitleCode: "b_uil",
	}
}
