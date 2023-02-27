package burkhan_buudai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比格尔BigerBarony struct {
	feud.BaseBarony
}

var BBiger比格尔 feud.Barony = &比格尔BigerBarony{}

func init() {
    f := BBiger比格尔.(*比格尔BigerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biger",
		TitleName: "比格尔",
		TitleCode: "b_biger",
	}
}
