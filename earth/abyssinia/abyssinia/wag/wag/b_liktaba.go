package wag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利克塔巴LiktabaBarony struct {
	feud.BaseBarony
}

var BLiktaba利克塔巴 feud.Barony = &利克塔巴LiktabaBarony{}

func init() {
    f := BLiktaba利克塔巴.(*利克塔巴LiktabaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liktaba",
		TitleName: "利克塔巴",
		TitleCode: "b_liktaba",
	}
}
