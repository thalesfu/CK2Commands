package hradec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫鲁迪姆ChrudimBarony struct {
	feud.BaseBarony
}

var BChrudim赫鲁迪姆 feud.Barony = &赫鲁迪姆ChrudimBarony{}

func init() {
    f := BChrudim赫鲁迪姆.(*赫鲁迪姆ChrudimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chrudim",
		TitleName: "赫鲁迪姆",
		TitleCode: "b_chrudim",
	}
}
